package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

	"github.com/andrescosta/ticketex/func/reservation/internal/handler"
	"github.com/andrescosta/ticketex/func/reservation/internal/repository"
)

type Config struct {
	CassandraHosts    string `json:"cassandra_hosts"`
	CassandraKeyspace string `json:"cassandra_keyspace"`
}

func main() {
	config := loadConfig()
	cluster := gocql.NewCluster(config.CassandraHosts)
	cluster.Keyspace = config.CassandraKeyspace
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra:", err)
	}
	defer session.Close()

	dataAccess := &repository.CassandraDataAccess{Session: session}

	// Create the router
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("reservation-service"))

	// Define the routes
	router.HandleFunc("/reservations", handler.GetReservationsHandler(dataAccess)).Methods(http.MethodGet)
	router.HandleFunc("/reservations", handler.CreateReservationHandler(dataAccess)).Methods(http.MethodPost)
	router.HandleFunc("/reservations/{id}", handler.PatchReservationHandler(dataAccess)).Methods(http.MethodPatch)

	log.Fatal(http.ListenAndServe(":8080", router))
	println("Started")
}

func loadConfig() Config {
	file, err := os.Open("./config.json")
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
