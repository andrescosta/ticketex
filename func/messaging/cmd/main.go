package main

import (
	"log"
	"net/http"

	"github.com/andrescosta/ticketex/func/messaging/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"

	"github.com/andrescosta/ticketex/func/messaging/internal/resource"
)

func main() {
	config := config.Load("../config.json")
	if msg, err := resource.Init(); err == nil {
		router := chi.NewRouter()
		router.Use(middleware.Logger)

		logger := httplog.NewLogger("messaging-log", httplog.Options{
			JSON: true,
		})

		router.Mount("/v1/messaging", msg.Routes(logger))
		log.Println("Server listening at", config.Host)
		log.Fatal(http.ListenAndServe(config.Host, router))
	} else {
		log.Fatal(err)
	}
}
