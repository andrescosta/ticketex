package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andrescosta/ticketex/func/reservation/internal/repository"
	"github.com/andrescosta/ticketex/func/reservation/internal/resource"
)

type Config struct {
	PostgressDsn string `json:"postgress_dsn"`
	Host         string `json:"host"`
}

func main() {
	config := loadConfig()
	dataAccess, err := repository.Init(config.PostgressDsn)
	if err != nil {
		log.Fatal(err)
	}
	reservation := resource.ReservationResource{DataAccess: dataAccess}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/reservations", reservation.Routes())
	log.Println("Server listening on", config.Host)
	log.Fatal(http.ListenAndServe(config.Host, router))
}

func loadConfig() Config {
	file, err := os.Open("../../config.json")
	if err != nil {
		log.Fatal("Failed to open config file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to decode config file:", err)
	}

	return config
}
