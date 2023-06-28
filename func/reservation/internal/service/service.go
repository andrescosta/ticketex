package service

import (
	"errors"

	"github.com/andrescosta/ticketex/func/reservation/internal/config"
	"github.com/andrescosta/ticketex/func/reservation/internal/entity"
	"github.com/andrescosta/ticketex/func/reservation/internal/enums"
	"github.com/andrescosta/ticketex/func/reservation/internal/repository"
	"github.com/andrescosta/ticketex/func/reservation/internal/rerrors"
	"gorm.io/gorm"
)

type IReservationSvc interface {
	NewReservationMetadata(reservation entity.ReservationMetadata) error
	GetMetadata(adventureId string) (entity.ReservationMetadata, error)
	AddMoreAvailability(reservationCapacity entity.ReservationCapacity) error
	NewReservationTypeMetadata(reservationCapacity entity.ReservationCapacity) error
	Reserve(reservation entity.Reservation) error
	Paid(reservation entity.Reservation) error
	Cancelled(reservation entity.Reservation) error
}

type ReservationSvc struct {
	repo repository.IReservation
}

func Init(config config.Config) (IReservationSvc, error) {
	if repo, err := repository.Init(config); err != nil {
		return nil, err
	} else {
		return &ReservationSvc{
			repo: repo,
		}, nil
	}
}

func (r *ReservationSvc) NewReservationMetadata(reservation entity.ReservationMetadata) error {
	return r.repo.AddReservationMetadata(reservation)
}

func (r *ReservationSvc) GetMetadata(adventureId string) (entity.ReservationMetadata, error) {
	return r.repo.GetReservationMetadata(adventureId)
}

func (r *ReservationSvc) AddMoreAvailability(reservationCapacity entity.ReservationCapacity) error {
	return r.repo.AddMoreAvailability(reservationCapacity)
}

func (r *ReservationSvc) NewReservationTypeMetadata(reservationCapacity entity.ReservationCapacity) error {
	return r.repo.AddReservationCapacity(reservationCapacity)
}

func (r *ReservationSvc) Reserve(reservation entity.Reservation) error {
	res, err := r.repo.GetReservationUnscoped(reservation)
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res = reservation
			res.Status = enums.Pending
			return r.repo.ReserveIfAvailableCapacity(reservation)
		} else {
			if res.DeletedAt.Valid {
				return rerrors.ErrIllegalReservationStatusDeleted
			}
			switch res.Status {
			case enums.Canceled:
				res.Status = enums.Pending
				res.Quantity = reservation.Quantity
				return r.repo.ReserveAndRecycleIfAvailableCapacity(res)
			case enums.Reserved:
				return rerrors.ErrIllegalReservationStatusReserved
			case enums.Pending:
				return rerrors.ErrIllegalReservationStatusPending
			default:
				return rerrors.ErrIllegalReservationStatus
			}
		}
	} else {
		return err
	}
}

func (r *ReservationSvc) Paid(reservation entity.Reservation) error {
	if res, err := r.repo.GetReservation(reservation); err == nil {
		switch res.Status {
		case enums.Canceled:
			return rerrors.ErrIllegalReservationStatusCanceled
		case enums.Reserved:
			return rerrors.ErrIllegalReservationStatusReserved
		case enums.Pending:
			res.Status = enums.Reserved
			return r.repo.UpdateReservation(res)
		default:
			return rerrors.ErrIllegalReservationStatus
		}
	} else {
		return err
	}
}

func (r *ReservationSvc) Cancelled(reservation entity.Reservation) error {
	if res, err := r.repo.GetReservation(reservation); err == nil {
		switch res.Status {
		case enums.Canceled:
			return rerrors.ErrIllegalReservationStatusCanceled
		case enums.Reserved:
			return rerrors.ErrIllegalReservationStatusReserved
		case enums.Pending:
			return r.repo.CancelReservation(res)
		default:
			return rerrors.ErrIllegalReservationStatus
		}
	} else {
		return err
	}
}
