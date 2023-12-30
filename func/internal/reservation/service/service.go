package service

import (
	"errors"

	"github.com/andrescosta/ticketex/func/internal/config"
	"github.com/andrescosta/ticketex/func/internal/reservation/entity"
	"github.com/andrescosta/ticketex/func/internal/reservation/enum"
	"github.com/andrescosta/ticketex/func/internal/reservation/repository"
	"gorm.io/gorm"
)

type Reservation struct {
	repo *repository.Reservation
}

func New(config config.Config) (*Reservation, error) {
	repo, err := repository.New(config)
	if err != nil {
		return nil, err
	}
	return &Reservation{
		repo: repo,
	}, nil
}

func (r *Reservation) Get(reservation entity.Reservation) (entity.Reservation, error) {
	return r.repo.GetReservation(reservation)
}

func (r *Reservation) NewReservationMetadata(reservation entity.ReservationMetadata) error {
	return r.repo.AddReservationMetadata(reservation)
}

func (r *Reservation) GetMetadata(adventureID string) (entity.ReservationMetadata, error) {
	return r.repo.GetReservationMetadata(adventureID)
}

func (r *Reservation) AddMoreAvailability(reservationCapacity entity.ReservationCapacity) error {
	return r.repo.AddMoreAvailability(reservationCapacity)
}

func (r *Reservation) NewReservationTypeMetadata(reservationCapacity entity.ReservationCapacity) error {
	return r.repo.AddReservationCapacity(reservationCapacity)
}

func (r *Reservation) Reserve(reservation entity.Reservation) error {
	res, err := r.repo.GetReservationUnscoped(reservation)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res = reservation
			res.Status = enum.Pending
			return r.repo.ReserveIfAvailableCapacity(reservation)
		}
		return err
	}
	if res.DeletedAt.Valid {
		return repository.ErrIllegalReservationStatusDeleted
	}
	switch res.Status {
	case enum.Canceled:
		res.Status = enum.Pending
		res.Quantity = reservation.Quantity
		return r.repo.ReserveAndRecycleIfAvailableCapacity(res)
	case enum.Reserved:
		return repository.ErrIllegalReservationStatusReserved
	case enum.Pending:
		return repository.ErrIllegalReservationStatusPending
	default:
		return repository.ErrIllegalReservationStatus
	}
}

func (r *Reservation) Paid(reservation entity.Reservation) error {
	if res, err := r.repo.GetReservation(reservation); err == nil {
		switch res.Status {
		case enum.Canceled:
			return repository.ErrIllegalReservationStatusCanceled
		case enum.Reserved:
			return repository.ErrIllegalReservationStatusReserved
		case enum.Pending:
			res.Status = enum.Reserved
			return r.repo.UpdateReservation(res)
		default:
			return repository.ErrIllegalReservationStatus
		}
	} else {
		return err
	}
}

func (r *Reservation) Cancelled(reservation entity.Reservation) error {
	if res, err := r.repo.GetReservation(reservation); err == nil {
		switch res.Status {
		case enum.Canceled:
			return repository.ErrIllegalReservationStatusCanceled
		case enum.Reserved:
			return repository.ErrIllegalReservationStatusReserved
		case enum.Pending:
			return r.repo.CancelReservation(res)
		default:
			return repository.ErrIllegalReservationStatus
		}
	} else {
		return err
	}
}
