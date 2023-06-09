package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"

	"github.com/andrescosta/ticketex/func/reservation/internal/config"
	"github.com/andrescosta/ticketex/func/reservation/internal/entity"
	"github.com/andrescosta/ticketex/func/reservation/internal/enums"
	"github.com/andrescosta/ticketex/func/reservation/internal/model"
	"github.com/andrescosta/ticketex/func/reservation/internal/rerrors"
	"github.com/andrescosta/ticketex/func/reservation/internal/service"
)

type IReservationResource interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	PatchCapacity(w http.ResponseWriter, r *http.Request)
	PostCapacity(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
	PatchUser(w http.ResponseWriter, r *http.Request)
	Routes(logger zerolog.Logger) chi.Router
}

type ReservationResource struct {
	service service.IReservationSvc
}

func Init(config config.Config) (IReservationResource, error) {

	svc, err := service.Init(config)
	if err != nil {
		return nil, err
	}
	reservation := &ReservationResource{
		service: svc,
	}

	return reservation, nil
}

func (rr ReservationResource) Routes(logger zerolog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Post("/metadata", rr.Post)

	r.Route("/{adventure_id}/capacities/{type}/users/{user_id}", func(r1 chi.Router) {
		r1.Post("/", rr.PostUser)
		r1.Get("/", rr.GetUser)
		r1.Patch("/status/{status}", rr.PatchUser)
	})

	r.Route("/metadata/{adventure_id}", func(r2 chi.Router) {
		r2.Get("/", rr.Get)
		r2.Route("/capacities", func(r3 chi.Router) {
			r3.Post("/", rr.PostCapacity)
			r3.Route("/{type}", func(r4 chi.Router) {
				r4.Patch("/", rr.PatchCapacity)
			})
		})
	})

	return r
}

func (rr ReservationResource) Get(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationMetadata{Adventure_id: chi.URLParam(r, "adventure_id")}
	if reservation, err := rr.getReservation(res, w); err != nil {
		rr.logError("Failed to get reservation:", err, r)
		http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
	} else {
		if err = json.NewEncoder(w).Encode(reservation); err != nil {
			rr.logError("Failed to get reservation:", err, r)
			http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (rr ReservationResource) getReservation(res entity.ReservationMetadata, w http.ResponseWriter) (model.ReservationMetadata, error) {
	if reservation, err := rr.service.GetMetadata(res.Adventure_id); err != nil {
		return model.ReservationMetadata{}, err
	} else {
		rreservation := model.ReservationMetadata{
			Adventure_id: reservation.Adventure_id,
			Status:       enums.ReservationMetadataStatus(reservation.Status),
		}
		var capacities []model.Capacity
		for _, v := range reservation.Capacities {
			capacity := model.Capacity{
				Type:         v.Type,
				Availability: v.Availability,
			}
			capacities = append(capacities, capacity)
		}
		rreservation.Capacity = capacities
		return rreservation, nil
	}
}

func (rr ReservationResource) Post(w http.ResponseWriter, r *http.Request) {
	var mReservation model.ReservationMetadata
	if err := json.NewDecoder(r.Body).Decode(&mReservation); err != nil {
		rr.logError("Failed to decode request body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	var newReservationCapacities []entity.ReservationCapacity
	var newReservation entity.ReservationMetadata
	newReservation.Adventure_id = mReservation.Adventure_id
	newReservation.Status = uint(mReservation.Status)
	for _, v := range mReservation.Capacity {
		var newReservationCapacity entity.ReservationCapacity
		newReservationCapacity.Adventure_id = mReservation.Adventure_id
		newReservationCapacity.Availability = v.Availability
		newReservationCapacity.Type = v.Type
		newReservationCapacities = append(newReservationCapacities, newReservationCapacity)
	}
	newReservation.Capacities = newReservationCapacities
	err := rr.service.NewReservationMetadata(newReservation)
	if err != nil {
		rr.logError("Failed to create reservation:", err, r)
		http.Error(w, "Failed to create reservation", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&mReservation)
}

func (rr ReservationResource) PatchCapacity(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationCapacity{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		Type:         chi.URLParam(r, "type"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		rr.logError("Failed to decode body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := rr.service.AddMoreAvailability(res)
	if err != nil {
		rr.logError("Error updating capacity:", err, r)
		if errors.Is(err, rerrors.ErrIllegalAvailability) {
			http.Error(w, "Illegal availability.", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rr ReservationResource) PostCapacity(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationCapacity{
		Adventure_id: chi.URLParam(r, "adventure_id"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		rr.logError("Failed to decode body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := rr.service.NewReservationTypeMetadata(res)
	if err != nil {
		rr.logError("Failed to create metadata:", err, r)
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (rr ReservationResource) GetUser(w http.ResponseWriter, r *http.Request) {
	rese, err := rr.service.Get(
		entity.Reservation{
			Adventure_id: chi.URLParam(r, "adventure_id"),
			Type:         chi.URLParam(r, "type"),
			User_id:      chi.URLParam(r, "user_id"),
		})

	if err != nil {
		rr.logError("Failed to get reservation:", err, r)
		http.Error(w, "Failed to get reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(
		&model.Reservation{
			Adventure_id: rese.Adventure_id,
			Type:         rese.Type,
			Quantity:     rese.Quantity,
			Status:       rese.Status,
			User_id:      rese.User_id,
		})
}

func (rr ReservationResource) PostUser(w http.ResponseWriter, r *http.Request) {
	res := entity.Reservation{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		Type:         chi.URLParam(r, "type"),
		User_id:      chi.URLParam(r, "user_id"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		rr.logError("Failed to decode reservation body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	res.Status = enums.Pending
	err := rr.service.Reserve(res)
	if err != nil {
		rr.logError("Failed to create reservation:", err, r)
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&res)
}

func (rr ReservationResource) PatchUser(w http.ResponseWriter, r *http.Request) {
	res := entity.Reservation{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		Type:         chi.URLParam(r, "type"),
		User_id:      chi.URLParam(r, "user_id"),
	}

	status := enums.ToReservationUserStatus(chi.URLParam(r, "status"))

	var err error

	switch status {
	case enums.Pending:
		rr.logError("Failed to update reservation", nil, r)
		http.Error(w, "Failed to update reservation", http.StatusBadRequest)
		return
	case enums.Reserved:
		err = rr.service.Paid(res)
		if err != nil {
			rr.logError("Failed to update reservation:", err, r)
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
			return
		}
	case enums.Canceled:
		err = rr.service.Cancelled(res)
		if err != nil {
			rr.logError("Failed to update reservation:", err, r)
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (rr ReservationResource) logError(msg string, err error, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	if err != nil {
		msg = fmt.Sprint(msg, err)
	}
	oplog.Error().Msg(msg)
}
