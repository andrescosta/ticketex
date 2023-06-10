package repository

import (
	"context"
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

func (d *CassandraDataAccess) GetReservations(id int) (model.Reservation, error) {
	iter := d.Session.Query("SELECT adventure_id,status,type,max_people,availability FROM reservations WHERE adventure_id=?", id).Iter()
	var reservation model.Reservation
	var capacity model.Capacity
	for iter.Scan(&reservation.Adventure_id, &reservation.Status, &capacity.Type, &capacity.Max, &capacity.Availability) {
		reservation.Capacity = append(reservation.Capacity, capacity)
	}
	if err := iter.Close(); err != nil {
		return reservation, err
	}

	return reservation, nil
}

func (d *CassandraDataAccess) CreateReservation(reservation model.Reservation) error {
	ctx := context.Background()

	b := d.Session.NewBatch(gocql.UnloggedBatch).WithContext(ctx)

	for _, v := range reservation.Capacity {
		b.Entries = append(b.Entries, gocql.BatchEntry{
			Stmt:       "INSERT INTO reservations (adventure_id,status,type,max_people,availability) VALUES (?, ?, ?, ?, ?)",
			Args:       []interface{}{reservation.Adventure_id, reservation.Status, v.Type, v.Max, v.Availability},
			Idempotent: true,
		})
	}
	return d.Session.ExecuteBatch(b)
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
