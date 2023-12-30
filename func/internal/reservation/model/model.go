package model

import "github.com/andrescosta/ticketex/func/internal/reservation/enums"

type ReservationMetadata struct {
	AdventureID string                          `json:"adventure_id"`
	Status      enums.ReservationMetadataStatus `json:"status"`
	Capacity    []Capacity                      `json:"capacities"`
}

type Capacity struct {
	Type         string `json:"type"`
	Availability uint   `json:"availability"`
}

type Reservation struct {
	AdventureID string                  `json:"adventure_id"`
	UserID      string                  `json:"user_id"`
	Type        string                  `json:"type"`
	Quantity    uint                    `json:"quantity"`
	Status      enums.ReservationStatus `json:"status"`
}
