package model

import "strings"

type ReservationStatus uint
type ReservationUserStatus uint

const (
	Closed ReservationStatus = iota
	Opened
	Autumn
	Winter
	Spring
)

const (
	Pending ReservationUserStatus = iota
	Reserved
	Canceled
)

func ToReservationUserStatus(s string) ReservationUserStatus {
	switch strings.ToLower(s) {
	case "pending":
		return Pending
	case "reserved":
		return Reserved
	case "canceled":
		return Canceled
	}
	return Pending
}

type Reservation struct {
	Adventure_id string            `json:"adventure_id"`
	Status       ReservationStatus `json:"status"`
	Capacity     []Capacity        `json:"capacities"`
}

type Capacity struct {
	Type         string `json:"type"`
	Availability uint   `json:"availability"`
	Max          uint   `json:"max"`
}

type ReservationUser struct {
	Adventure_id string                `json:"adventure_id"`
	User_id      string                `json:"user_id"`
	Type         string                `json:"type"`
	Quantity     string                `json:"quantity"`
	Status       ReservationUserStatus `json:"status"`
}
