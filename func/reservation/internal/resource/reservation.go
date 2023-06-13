package resource

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andrescosta/ticketex/func/reservation/internal/entity"
	"github.com/andrescosta/ticketex/func/reservation/internal/model"
	"github.com/andrescosta/ticketex/func/reservation/internal/repository"
)

type ReservationResource struct {
	DataAccess repository.DataAccess
}

func (rr ReservationResource) Routes() chi.Router {
	rtr := chi.NewRouter()
	rtr.Post("/", rr.Post)

	rtr.Route("/{adventure_id}", func(rrtr chi.Router) {
		rrtr.Get("/", rr.Get)
		rrtr.Route("/capacities", func(rrrtr chi.Router) {
			rrrtr.Post("/", rr.PostCapacity)
			rrrtr.Route("/{type}", func(rrrrtr chi.Router) {
				rrrrtr.Patch("/", rr.PatchCapacity)
				rrrrtr.Route("/users/{user_id}", func(rrrrrtr chi.Router) {
					rrrrrtr.Post("/", rr.PostUser)
					rrrrrtr.Patch("/status/{status}", rr.PatchUser)
				})
			})
		})
	})

	return rtr
}

func (rr ReservationResource) Get(w http.ResponseWriter, r *http.Request) {
	res := entity.Reservation{Adventure_id: chi.URLParam(r, "adventure_id")}
	if reservation, err := rr.getReservation(res, w); err != nil {
		http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
	} else {
		if err = json.NewEncoder(w).Encode(reservation); err != nil {
			http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (rr ReservationResource) getReservation(res entity.Reservation, w http.ResponseWriter) (model.Reservation, error) {
	if reservation, reservationc, err := rr.DataAccess.GetReservation(res); err != nil {
		return model.Reservation{}, err
	} else {
		rreservation := model.Reservation{
			Adventure_id: reservation.Adventure_id,
			Status:       model.ReservationStatus(reservation.Status),
		}
		var capacities []model.Capacity
		for _, v := range reservationc {
			capacity := model.Capacity{
				Type:         v.Type,
				Availability: v.Availability,
				Max:          v.Max,
			}
			capacities = append(capacities, capacity)
		}
		rreservation.Capacity = capacities
		return rreservation, nil
	}
}

func (rr ReservationResource) Post(w http.ResponseWriter, r *http.Request) {
	var mReservation model.Reservation
	if err := json.NewDecoder(r.Body).Decode(&mReservation); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	var newReservationCapacities []entity.ReservationCapacity
	var newReservation entity.Reservation
	newReservation.Adventure_id = mReservation.Adventure_id
	newReservation.Status = uint(mReservation.Status)
	for _, v := range mReservation.Capacity {
		var newReservationCapacity entity.ReservationCapacity
		newReservationCapacity.Adventure_id = mReservation.Adventure_id
		newReservationCapacity.Availability = v.Availability
		newReservationCapacity.Max = v.Max
		newReservationCapacity.Type = v.Type
		newReservationCapacities = append(newReservationCapacities, newReservationCapacity)
	}
	err := rr.DataAccess.CreateReservations(newReservation, newReservationCapacities)
	if err != nil {
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
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := rr.DataAccess.PatchReservationCapacity(res)
	if err != nil {
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rr ReservationResource) PostCapacity(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationCapacity{
		Adventure_id: chi.URLParam(r, "adventure_id"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := rr.DataAccess.PostReservationCapacity(res)
	if err != nil {
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rr ReservationResource) PostUser(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationUser{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		Type:         chi.URLParam(r, "type"),
		User_id:      chi.URLParam(r, "user_id"),
	}

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	res.Status = uint(model.Pending)
	err := rr.DataAccess.PostReservationUser(res)
	if err != nil {
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&res)
}

func (rr ReservationResource) PatchUser(w http.ResponseWriter, r *http.Request) {
	res := entity.ReservationUser{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		Type:         chi.URLParam(r, "type"),
		User_id:      chi.URLParam(r, "user_id"),
		Status:       uint(model.ToReservationUserStatus(chi.URLParam(r, "status"))),
	}

	err := rr.DataAccess.PatchReservationUser(res)
	if err != nil {
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
