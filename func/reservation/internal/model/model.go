package model

import "github.com/andrescosta/ticketex/func/reservation/internal/enums"

type Reservation struct {
	Adventure_id string                  `json:"adventure_id"`
	Status       enums.ReservationStatus `json:"status"`
	Capacity     []Capacity              `json:"capacities"`
}

type Capacity struct {
	Type         string `json:"type"`
	Availability uint   `json:"availability"`
}

type ReservationUser struct {
	Adventure_id string                      `json:"adventure_id"`
	User_id      string                      `json:"user_id"`
	Type         string                      `json:"type"`
	Quantity     string                      `json:"quantity"`
	Status       enums.ReservationUserStatus `json:"status"`
}
