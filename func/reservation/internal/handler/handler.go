package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrescosta/ticketex/func/reservation/internal/model"
	"github.com/andrescosta/ticketex/func/reservation/internal/repository"
	"go.opentelemetry.io/otel"
)

func GetReservationsHandler(dataAccess repository.DataAccess) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("reservation-service")
		_, span := tracer.Start(ctx, "getReservationsHandler")
		defer span.End()

		reservations, err := dataAccess.GetReservations()
		if err != nil {
			otel.Handle(fmt.Errorf("Failed to get reservations: %v", err))
			http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(reservations)
		if err != nil {
			otel.Handle(fmt.Errorf("Failed to marshal response: %v", err))
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func CreateReservationHandler(dataAccess repository.DataAccess) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("reservation-service")
		_, span := tracer.Start(ctx, "createReservationHandler")
		defer span.End()

		var newReservation model.Reservation
		if err := json.NewDecoder(r.Body).Decode(&newReservation); err != nil {
			otel.Handle(fmt.Errorf("Failed to decode request body: %v", err))
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		err := dataAccess.CreateReservation(newReservation)
		if err != nil {
			otel.Handle(fmt.Errorf("Failed to create reservation: %v", err))
			http.Error(w, "Failed to create reservation", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func PatchReservationHandler(dataAccess repository.DataAccess) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("reservation-service")
		_, span := tracer.Start(ctx, "patchReservationHandler")
		defer span.End()

		var updatedReservation model.Reservation
		if err := json.NewDecoder(r.Body).Decode(&updatedReservation); err != nil {
			otel.Handle(fmt.Errorf("Failed to decode request body: %v", err))
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		err := dataAccess.PatchReservation(updatedReservation)
		if err != nil {
			otel.Handle(fmt.Errorf("Failed to update reservation: %v", err))
			http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
