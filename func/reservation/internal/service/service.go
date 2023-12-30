package service

import (
	"errors"

	"github.com/andrescosta/ticketex/func/reservation/internal/config"
	"github.com/andrescosta/ticketex/func/reservation/internal/entity"
	"github.com/andrescosta/ticketex/func/reservation/internal/enums"
	"github.com/andrescosta/ticketex/func/reservation/internal/repository"
	"gorm.io/gorm"
)

type Reservation struct {
	repo *repository.Reservation
}

func New(config config.Config) (*Reservation, error) {
	if repo, err := repository.New(config); err != nil {
		return nil, err
	} else {
		return &Reservation{
			repo: repo,
		}, nil
	}
}

func (r *Reservation) Get(reservation entity.Reservation) (entity.Reservation, error) {
	return r.repo.GetReservation(reservation)
}

func (r *Reservation) NewReservationMetadata(reservation entity.ReservationMetadata) error {
	return r.repo.AddReservationMetadata(reservation)
}

func (r *Reservation) GetMetadata(adventureId string) (entity.ReservationMetadata, error) {
	return r.repo.GetReservationMetadata(adventureId)
}

func (r *Reservation) AddMoreAvailability(reservationCapacity entity.ReservationCapacity) error {
	return r.repo.AddMoreAvailability(reservationCapacity)
}

func (r *Reservation) NewReservationTypeMetadata(reservationCapacity entity.ReservationCapacity) error {
	return r.repo.AddReservationCapacity(reservationCapacity)
}

func (r *Reservation) Reserve(reservation entity.Reservation) error {
	res, err := r.repo.GetReservationUnscoped(reservation)
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res = reservation
			res.Status = enums.Pending
			return r.repo.ReserveIfAvailableCapacity(reservation)
		} else {
			if res.DeletedAt.Valid {
				return repository.ErrIllegalReservationStatusDeleted
			}
			switch res.Status {
			case enums.Canceled:
				res.Status = enums.Pending
				res.Quantity = reservation.Quantity
				return r.repo.ReserveAndRecycleIfAvailableCapacity(res)
			case enums.Reserved:
				return repository.ErrIllegalReservationStatusReserved
			case enums.Pending:
				return repository.ErrIllegalReservationStatusPending
			default:
				return repository.ErrIllegalReservationStatus
			}
		}
	} else {
		return err
	}
}

func (r *Reservation) Paid(reservation entity.Reservation) error {
	if res, err := r.repo.GetReservation(reservation); err == nil {
		switch res.Status {
		case enums.Canceled:
			return repository.ErrIllegalReservationStatusCanceled
		case enums.Reserved:
			return repository.ErrIllegalReservationStatusReserved
		case enums.Pending:
			res.Status = enums.Reserved
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
		case enums.Canceled:
			return repository.ErrIllegalReservationStatusCanceled
		case enums.Reserved:
			return repository.ErrIllegalReservationStatusReserved
		case enums.Pending:
			return r.repo.CancelReservation(res)
		default:
			return repository.ErrIllegalReservationStatus
		}
	} else {
		return err
	}
}
