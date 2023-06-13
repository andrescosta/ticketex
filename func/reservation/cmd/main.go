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
}

func main() {
	config := loadConfig()
	dataAccess := &repository.PostgressDataAccess{}
	dataAccess.Init(config.PostgressDsn)
	reservation := resource.ReservationResource{DataAccess: dataAccess}

	// Create the router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/reservations", reservation.Routes())

	// Define the routes

	log.Fatal(http.ListenAndServe("localhost:8080", router))
	println("Started")
}

func loadConfig() Config {
	file, err := os.Open("../config.json")
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
