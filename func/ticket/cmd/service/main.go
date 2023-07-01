package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"

	"github.com/andrescosta/ticketex/func/ticket/internal/config"
	"github.com/andrescosta/ticketex/func/ticket/internal/resource"
)

func main() {
	config := config.Load("../../config.json")
	if ticket, err := resource.Init(config); err == nil {
		router := chi.NewRouter()
		router.Use(middleware.Logger)

		logger := httplog.NewLogger("ticket-log", httplog.Options{
			JSON: true,
		})

		router.Mount("/v1/tickets", ticket.Routes(logger))
		log.Println("Server listening at", config.Host)
		log.Fatal(http.ListenAndServe(config.Host, router))
	} else {
		log.Fatal(err)
	}
}
