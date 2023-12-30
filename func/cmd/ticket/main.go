package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"

	"github.com/andrescosta/ticketex/func/internal/config"
	"github.com/andrescosta/ticketex/func/internal/ticket/resource"
)

func main() {
	config := config.Load("../../configs/config-t.json")
	ticket, err := resource.New(config)
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	logger := httplog.NewLogger("ticket-log", httplog.Options{
		JSON: true,
	})

	router.Mount("/tickets", ticket.Routes(logger))
	log.Println("Server listening at", config.Host)
	listener, err := net.Listen("tcp", config.Host)
	if err != nil {
		log.Fatalf("error:%v", err)
	}

	server := &http.Server{
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("http.Serve: failed to serve: %v", err)
	}
}
