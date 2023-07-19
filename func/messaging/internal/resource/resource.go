package resource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrescosta/ticketex/func/messaging/internal/model"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

type IMessagingResource interface {
	Post(w http.ResponseWriter, r *http.Request)
	Routes(logger zerolog.Logger) chi.Router
}

type MessagingResource struct {
}

func Init() (IMessagingResource, error) {

	reservation := &MessagingResource{}

	return reservation, nil
}

func (rr MessagingResource) Routes(logger zerolog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Route("/", func(r2 chi.Router) {
		r2.Post("/", rr.Post)
	})
	return r
}

func (rr MessagingResource) Post(w http.ResponseWriter, r *http.Request) {
	var msg model.Message
	userid, ok := rr.getUserId(r, w)
	if !ok {
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		rr.logError("Failed to decode request body:", err, r)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	oplog := httplog.LogEntry(r.Context())
	msg1 := fmt.Sprintf("USer: %s Message: %s", userid, msg)
	oplog.Info().Msg(msg1)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (rr MessagingResource) logError(msg string, err error, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	if err != nil {
		msg = fmt.Sprint(msg, err)
	}
	oplog.Error().Msg(msg)
}
func (rr MessagingResource) getUserId(r *http.Request, w http.ResponseWriter) (string, bool) {
	claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		rr.logError("Failed to validate token", nil, r)
		http.Error(w, "Failed to validate token", http.StatusUnauthorized)
		return "", true
	}
	userid := claims.RegisteredClaims.Subject
	return userid, false
}
