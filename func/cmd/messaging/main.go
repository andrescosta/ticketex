package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/andrescosta/ticketex/func/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"

	"github.com/andrescosta/ticketex/func/internal/messaging/resource"
)

func main() {
	config := config.Load("../../configs/config-m.json")
	msg, err := resource.New()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	logger := httplog.NewLogger("messaging-log", httplog.Options{
		JSON: true,
	})

	router.Mount("/messaging", msg.Routes(logger))
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
