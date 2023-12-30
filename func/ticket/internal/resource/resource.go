package resource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"

	"github.com/andrescosta/ticketex/func/ticket/internal/config"
	"github.com/andrescosta/ticketex/func/ticket/internal/entity"
	"github.com/andrescosta/ticketex/func/ticket/internal/model"
	"github.com/andrescosta/ticketex/func/ticket/internal/service"
)

type Ticket struct {
	service *service.Ticket
}

func New(config config.Config) (*Ticket, error) {

	svc, err := service.New(config)
	if err != nil {
		return nil, err
	}
	reservation := &Ticket{
		service: svc,
	}

	return reservation, nil
}

func (rr Ticket) Routes(logger zerolog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Route("/{adventure_id}/{type}/{user_id}", func(r2 chi.Router) {
		r2.Post("/", rr.Post)
		r2.Get("/", rr.Get)
	})
	return r
}

func (rr Ticket) Get(w http.ResponseWriter, r *http.Request) {
	tt := entity.TicketTrans{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		User_id:      chi.URLParam(r, "user_id"),
		Type:         chi.URLParam(r, "type")}
	if ticketTrans, err := rr.service.GetTicketTrans(tt); err != nil {
		rr.logError("Failed to get reservation:", err, r)
		http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
	} else {
		tr := rr.buildTrans(ticketTrans)
		if err = json.NewEncoder(w).Encode(&tr); err != nil {
			rr.logError("Failed to get reservation:", err, r)
			http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (rr Ticket) Post(w http.ResponseWriter, r *http.Request) {
	tt := entity.TicketTrans{
		Adventure_id: chi.URLParam(r, "adventure_id"),
		User_id:      chi.URLParam(r, "user_id"),
		Type:         chi.URLParam(r, "type"),
	}

	if err := json.NewDecoder(r.Body).Decode(&tt); err != nil {
		rr.logError("Failed to decode request body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	tickets, err := rr.service.GenerateTickets(tt)
	if err != nil {
		rr.logError("Failed to create reservation:", err, r)
		http.Error(w, "Failed to create reservation", http.StatusInternalServerError)
		return
	}
	res := rr.buildTrans(tickets)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&res)
}

func (Ticket) buildTrans(tt entity.TicketTrans) model.TicketTrans {
	res := model.TicketTrans{
		Adventure_id: tt.Adventure_id,
		User_id:      tt.User_id,
		Type:         tt.Type,
		Quantity:     tt.Quantity,
		Status:       tt.Status,
	}
	rest := make([]model.Ticket, len(tt.Tickets))
	for i, v := range tt.Tickets {
		rest[i] = model.Ticket{
			Code: v.Code,
		}
	}
	res.Tickets = rest
	return res
}

func (rr Ticket) logError(msg string, err error, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	if err != nil {
		msg = fmt.Sprint(msg, err)
	}
	oplog.Error().Msg(msg)
}
