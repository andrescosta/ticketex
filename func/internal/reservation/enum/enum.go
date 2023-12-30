package enum

import "strings"

type (
	ReservationMetadataStatus uint
	ReservationStatus         uint
)

const (
	Closed ReservationMetadataStatus = iota
	Opened
)

const (
	Pending ReservationStatus = iota
	Reserved
	Canceled
)

func ToReservationMetadataStatus(s string) ReservationMetadataStatus {
	switch strings.ToLower(s) {
	case "closed":
		return Closed
	case "opened":
		return Opened
	}
	return Opened
}

func ToReservationUserStatus(s string) ReservationStatus {
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
