package resource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrescosta/ticketex/func/messaging/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

type Messaging struct {
}

func New() (*Messaging, error) {

	reservation := &Messaging{}

	return reservation, nil
}

func (m Messaging) Routes(logger zerolog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Route("/{user_id}", func(r2 chi.Router) {
		r2.Post("/", m.Post)
	})
	return r
}

func (m Messaging) Post(w http.ResponseWriter, r *http.Request) {
	var msg model.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		m.logError("Failed to decode request body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	oplog := httplog.LogEntry(r.Context())
	msg1 := fmt.Sprint("Message:", msg)
	oplog.Info().Msg(msg1)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (m Messaging) logError(msg string, err error, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	if err != nil {
		msg = fmt.Sprint(msg, err)
	}
	oplog.Error().Msg(msg)
}
