package repository

import (
	"fmt"
	"strings"

	"github.com/andrescosta/ticketex/func/reservation/internal/model"
	"github.com/gocql/gocql"
)

type Config struct {
	CassandraHosts    string `json:"cassandra_hosts"`
	CassandraKeyspace string `json:"cassandra_keyspace"`
}

type DataAccess interface {
	GetReservations() ([]model.Reservation, error)
	CreateReservation(reservation model.Reservation) error
	PatchReservation(reservation model.Reservation) error
}

type CassandraDataAccess struct {
	Session *gocql.Session
}

func (d *CassandraDataAccess) GetReservations() ([]model.Reservation, error) {
	var reservations []model.Reservation

	iter := d.Session.Query("SELECT * FROM reservations").Iter()
	var reservation model.Reservation
	for iter.Scan(&reservation.ID, &reservation.Adventure.ID, &reservation.Capacity.Type, &reservation.Capacity.Current, &reservation.Capacity.Max) {
		reservations = append(reservations, reservation)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (d *CassandraDataAccess) CreateReservation(reservation model.Reservation) error {
	query := d.Session.Query("INSERT INTO reservations (id, adventure, capacity) VALUES (?, ?, ?)",
		reservation.ID, reservation.Adventure.ID, reservation.Capacity)
	return query.Exec()
}

func (d *CassandraDataAccess) PatchReservation(reservation model.Reservation) error {
	var fields []string
	var args []interface{}

	if reservation.Adventure.ID != "" {
		fields = append(fields, "adventure = ?")
		args = append(args, reservation.Adventure.ID)
	}
	if reservation.Capacity.Type != "" {
		fields = append(fields, "capacity = ?")
		args = append(args, reservation.Capacity)
	}

	queryString := fmt.Sprintf("UPDATE reservations SET %s WHERE id = ?", strings.Join(fields, ", "))
	args = append(args, reservation.ID)
	query := d.Session.Query(queryString, args...)

	return query.Exec()
}
