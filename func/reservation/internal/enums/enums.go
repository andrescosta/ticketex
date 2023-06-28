package enums

import "strings"

type ReservationStatus uint
type ReservationUserStatus uint

const (
	Closed ReservationStatus = iota
	Opened
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
