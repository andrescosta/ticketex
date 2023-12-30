package repository

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/andrescosta/ticketex/func/reservation/internal/config"
	"github.com/andrescosta/ticketex/func/reservation/internal/entity"
	"github.com/andrescosta/ticketex/func/reservation/internal/enums"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrOutOfCapacity                    = errors.New("no room")
	ErrUpdateCapacity                   = errors.New("capacity not found")
	ErrUpdateReservation                = errors.New("reservation not found")
	ErrIllegalAvailability              = errors.New("illegal availability")
	ErrIllegalReservationStatusCanceled = errors.New("illegal reservation status: canceled")
	ErrIllegalReservationStatusReserved = errors.New("illegal reservation status: reserved")
	ErrIllegalReservationStatusDeleted  = errors.New("illegal reservation status: deleted")
	ErrIllegalReservationStatusPending  = errors.New("illegal reservation status: pending")
	ErrIllegalReservationStatus         = errors.New("illegal reservation status: ")
)

type Reservation struct {
	DB *gorm.DB
}

func New(config config.Config) (*Reservation, error) {
	dataAccess := &Reservation{}
	loglevel := logger.Error
	if config.DebugSql {
		loglevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  loglevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(config.PostgressDsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		print(err)
		return nil, err
	}
	dataAccess.DB = db
	return dataAccess, nil
}

func (d *Reservation) GetReservationMetadata(adventureId string) (entity.ReservationMetadata, error) {
	metadata := entity.ReservationMetadata{Adventure_id: adventureId}
	res := d.DB.Model(&entity.ReservationMetadata{}).Preload("Capacities").First(&metadata)
	return metadata, res.Error
}
func (d *Reservation) GetReservation(reservation entity.Reservation) (entity.Reservation, error) {
	return d.getReservation(reservation, false)
}
func (d *Reservation) GetReservationUnscoped(reservation entity.Reservation) (entity.Reservation, error) {
	return d.getReservation(reservation, true)
}
func (d *Reservation) getReservation(reservation entity.Reservation, unscoped bool) (entity.Reservation, error) {
	reserv := entity.Reservation{
		Adventure_id: reservation.Adventure_id,
		User_id:      reservation.User_id,
		Type:         reservation.Type,
	}
	var res *gorm.DB
	if unscoped {
		res = d.DB.Unscoped().First(&reserv)
	} else {
		res = d.DB.First(&reserv)
	}
	return reserv, res.Error
}

func (d *Reservation) AddReservationMetadata(reservation entity.ReservationMetadata) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		if res := tx.Create(&reservation); res.Error != nil {
			return res.Error
		}
		return nil
	})
}

func (d *Reservation) UpdateReservationMetadata(reservation entity.ReservationMetadata) error {
	result := d.DB.Updates(&reservation)
	return result.Error
}
func (d *Reservation) ReserveAndRecycleIfAvailableCapacity(reservationUser entity.Reservation) error {
	return d.reserveIfAvailableCapacity(reservationUser, true)
}
func (d *Reservation) ReserveIfAvailableCapacity(reservationUser entity.Reservation) error {
	return d.reserveIfAvailableCapacity(reservationUser, false)
}
func (d *Reservation) reserveIfAvailableCapacity(reservationUser entity.Reservation, update bool) error {
	tx := d.DB.Begin(&sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var availability uint
	result := tx.Raw(`SELECT availability FROM reservation_capacities 
						WHERE availability>0 and availability>=? and adventure_id=? and 
							  type=? and deleted_at is null FOR UPDATE;`,
		// TODO: check https://stackoverflow.com/questions/75761088/in-postgres-is-there-a-need-to-lock-a-row-in-a-table-using-for-update-if-the-qu#:~:text=Database%20locks%20are%20handled%20by,avoided%2C%20and%20locks%20handled%20properly.
		reservationUser.Quantity, reservationUser.Adventure_id, reservationUser.Type).Scan(&availability)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrOutOfCapacity
	}
	availability = availability - reservationUser.Quantity
	result = tx.Exec(`UPDATE reservation_capacities 
						SET availability=? 
						WHERE adventure_id=? and type=? and deleted_at is null`,
		availability, reservationUser.Adventure_id, reservationUser.Type)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrUpdateCapacity
	}
	if update {
		reservationUser.DeletedAt = gorm.DeletedAt{Valid: false}
		result = tx.Unscoped().Updates(&reservationUser)
	} else {
		result = tx.Create(&reservationUser)
	}
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func (d *Reservation) UpdateReservationCapacity(reservationCapacity entity.ReservationCapacity) error {
	result := d.DB.Updates(&reservationCapacity)
	return result.Error
}

func (d *Reservation) AddReservationCapacity(reservationCapacity entity.ReservationCapacity) error {
	result := d.DB.Create(&reservationCapacity)
	return result.Error
}

func (d *Reservation) CreateReservation(reservationUser entity.Reservation) error {
	result := d.DB.Create(&reservationUser)
	return result.Error
}

func (d *Reservation) UpdateReservation(reservationUser entity.Reservation) error {
	result := d.DB.Updates(&reservationUser)
	return result.Error
}

func (d *Reservation) CancelReservation(reservationUser entity.Reservation) error {
	tx := d.DB.Begin(&sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var availability uint
	result := tx.Raw(`SELECT availability FROM reservation_capacities 
						WHERE adventure_id=? and 
							  type=? and deleted_at is null FOR UPDATE;`,
		reservationUser.Adventure_id, reservationUser.Type).Scan(&availability)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrUpdateCapacity
	}
	availability = availability + reservationUser.Quantity
	result = tx.Exec(`UPDATE reservation_capacities 
						SET availability=? 
						WHERE adventure_id=? and type=? and deleted_at is null`,
		availability, reservationUser.Adventure_id, reservationUser.Type)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrUpdateCapacity
	}
	reservationUser.Status = enums.Canceled
	result = tx.Updates(&reservationUser)
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrUpdateReservation
	}
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func (d *Reservation) AddMoreAvailability(reservationCapacity entity.ReservationCapacity) error {
	tx := d.DB.Begin(&sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var availability int
	result := tx.Raw(`SELECT availability FROM reservation_capacities 
						WHERE adventure_id=? and 
							  type=? and deleted_at is null FOR UPDATE;`,
		reservationCapacity.Adventure_id, reservationCapacity.Type).Scan(&availability)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrIllegalAvailability
	}
	if reservationCapacity.Availability < uint(availability) {
		tx.Rollback()
		return ErrIllegalAvailability
	}
	result = tx.Exec(`UPDATE reservation_capacities 
						SET availability=? 
						WHERE adventure_id=? and type=? and deleted_at is null;`,
		reservationCapacity.Availability, reservationCapacity.Adventure_id, reservationCapacity.Type)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return ErrUpdateCapacity
	}
	return tx.Commit().Error
}
