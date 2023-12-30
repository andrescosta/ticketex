package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"

	"github.com/andrescosta/ticketex/func/internal/config"
	"github.com/andrescosta/ticketex/func/internal/reservation/entity"
	"github.com/andrescosta/ticketex/func/internal/reservation/enum"
	"github.com/andrescosta/ticketex/func/internal/reservation/model"
	"github.com/andrescosta/ticketex/func/internal/reservation/repository"
	"github.com/andrescosta/ticketex/func/internal/reservation/service"
)

type Reservation struct {
	service *service.Reservation
}

func New(config config.Config) (*Reservation, error) {
	svc, err := service.New(config)
	if err != nil {
		return nil, err
	}
	reservation := &Reservation{
		service: svc,
	}

	return reservation, nil
}

func (rr Reservation) Routes(logger zerolog.Logger) chi.Router {
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

func (rr Reservation) Get(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationMetadata{AdventureID: chi.URLParam(r, "adventure_id")}
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

func (rr Reservation) getReservation(res entity.ReservationMetadata, _ http.ResponseWriter) (model.ReservationMetadata, error) {
	reservation, err := rr.service.GetMetadata(res.AdventureID)
	if err != nil {
		return model.ReservationMetadata{}, err
	}
	rreservation := model.ReservationMetadata{
		AdventureID: reservation.AdventureID,
		Status:      enum.ReservationMetadataStatus(reservation.Status),
	}
	capacities := make([]model.Capacity, len(reservation.Capacities))
	for idx, v := range reservation.Capacities {
		capacity := model.Capacity{
			Type:         v.Type,
			Availability: v.Availability,
		}
		capacities[idx] = capacity
	}
	rreservation.Capacity = capacities
	return rreservation, nil
}

func (rr Reservation) Post(w http.ResponseWriter, r *http.Request) {
	var mReservation model.ReservationMetadata
	if err := json.NewDecoder(r.Body).Decode(&mReservation); err != nil {
		rr.logError("Failed to decode request body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	var newReservation entity.ReservationMetadata
	newReservation.AdventureID = mReservation.AdventureID
	newReservation.Status = uint(mReservation.Status)
	newReservationCapacities := make([]entity.ReservationCapacity, len(mReservation.Capacity))
	for idx, v := range mReservation.Capacity {
		var newReservationCapacity entity.ReservationCapacity
		newReservationCapacity.AdventureID = mReservation.AdventureID
		newReservationCapacity.Availability = v.Availability
		newReservationCapacity.Type = v.Type
		newReservationCapacities[idx] = newReservationCapacity
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

func (rr Reservation) PatchCapacity(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationCapacity{
		AdventureID: chi.URLParam(r, "adventure_id"),
		Type:        chi.URLParam(r, "type"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		rr.logError("Failed to decode body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := rr.service.AddMoreAvailability(res)
	if err != nil {
		rr.logError("Error updating capacity:", err, r)
		if errors.Is(err, repository.ErrIllegalAvailability) {
			http.Error(w, "Illegal availability.", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rr Reservation) PostCapacity(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationCapacity{
		AdventureID: chi.URLParam(r, "adventure_id"),
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

func (rr Reservation) GetUser(w http.ResponseWriter, r *http.Request) {
	rese, err := rr.service.Get(
		entity.Reservation{
			AdventureID: chi.URLParam(r, "adventure_id"),
			Type:        chi.URLParam(r, "type"),
			UserID:      chi.URLParam(r, "user_id"),
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
			AdventureID: rese.AdventureID,
			Type:        rese.Type,
			Quantity:    rese.Quantity,
			Status:      rese.Status,
			UserID:      rese.UserID,
		})
}

func (rr Reservation) PostUser(w http.ResponseWriter, r *http.Request) {
	res := entity.Reservation{
		AdventureID: chi.URLParam(r, "adventure_id"),
		Type:        chi.URLParam(r, "type"),
		UserID:      chi.URLParam(r, "user_id"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		rr.logError("Failed to decode reservation body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	res.Status = enum.Pending
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

func (rr Reservation) PatchUser(w http.ResponseWriter, r *http.Request) {
	res := entity.Reservation{
		AdventureID: chi.URLParam(r, "adventure_id"),
		Type:        chi.URLParam(r, "type"),
		UserID:      chi.URLParam(r, "user_id"),
	}

	status := enum.ToReservationUserStatus(chi.URLParam(r, "status"))

	var err error

	switch status {
	case enum.Pending:
		rr.logError("Failed to update reservation", nil, r)
		http.Error(w, "Failed to update reservation", http.StatusBadRequest)
		return
	case enum.Reserved:
		err = rr.service.Paid(res)
		if err != nil {
			rr.logError("Failed to update reservation:", err, r)
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
			return
		}
	case enum.Canceled:
		err = rr.service.Cancelled(res)
		if err != nil {
			rr.logError("Failed to update reservation:", err, r)
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (rr Reservation) logError(msg string, err error, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	if err != nil {
		msg = fmt.Sprint(msg, err)
	}
	oplog.Error().Msg(msg)
}
