package model

type ReservationStatus int64
type ReservationUserStatus int64

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

type Reservation struct {
	Adventure_id string            `json:"adventure_id"`
	Status       ReservationStatus `json:"status"`
	Capacity     []Capacity        `json:"capacities"`
}

type Capacity struct {
	Type         string `json:"type"`
	Availability int    `json:"availability"`
	Max          int    `json:"max"`
}

type ReservationUser struct {
	Adventure_id string                `json:"adventure_id"`
	User_id      string                `json:"user_id"`
	Type         string                `json:"type"`
	Quantity     string                `json:"quantity"`
	Status       ReservationUserStatus `json:"status"`
}
