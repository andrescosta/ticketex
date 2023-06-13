package repository

import (
	"log"
	"os"
	"time"

	"github.com/andrescosta/ticketex/func/reservation/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataAccess interface {
	Init(dsn string) error
	GetReservation(reservation entity.Reservation) (entity.Reservation, []entity.ReservationCapacity, error)
	CreateReservations(reservation entity.Reservation,
		reservationCapacities []entity.ReservationCapacity) error
	PatchReservation(reservation entity.Reservation) error
	PatchReservationCapacity(reservationCapacity entity.ReservationCapacity) error
	PostReservationCapacity(reservationCapacity entity.ReservationCapacity) error
	PostReservationUser(reservationCapacity entity.ReservationUser) error
	PatchReservationUser(reservationUser entity.ReservationUser) error
}

type PostgressDataAccess struct {
	DB *gorm.DB
}

func (d *PostgressDataAccess) Init(dsn string) error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		print(err)
		return err
	}
	d.DB = db
	println("creating schema ...")
	err = db.AutoMigrate(&entity.Reservation{}, &entity.ReservationCapacity{}, &entity.ReservationUser{})
	if err != nil {
		print(err)
		return err
	}
	println("end creating schema ...")
	return nil
}

func (d *PostgressDataAccess) GetReservation(reservation entity.Reservation) (entity.Reservation, []entity.ReservationCapacity, error) {
	var reservationq = entity.Reservation{Adventure_id: reservation.Adventure_id}
	var reservationr entity.Reservation
	var reservationc []entity.ReservationCapacity
	if err := d.DB.First(&reservationr, reservationq); err.Error != nil {
		return reservationr, reservationc, err.Error
	}

	if err := d.DB.Find(&reservationc, reservationq); err.Error != nil {
		return reservationr, reservationc, err.Error
	}
	return reservationr, reservationc, nil
}

func (d *PostgressDataAccess) CreateReservations(reservation entity.Reservation, reservationCapacities []entity.ReservationCapacity) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		if res := tx.Create(&reservation); res.Error != nil {
			return res.Error
		}
		for _, reservationCapacity := range reservationCapacities {
			if res := tx.Create(&reservationCapacity); res.Error != nil {
				return res.Error
			}
		}
		return nil
	})
}

func (d *PostgressDataAccess) PatchReservation(reservation entity.Reservation) error {
	result := d.DB.Updates(&reservation)
	return result.Error
}

func (d *PostgressDataAccess) PatchReservationCapacity(reservationCapacity entity.ReservationCapacity) error {
	result := d.DB.Updates(&reservationCapacity)
	return result.Error
}

func (d *PostgressDataAccess) PostReservationCapacity(reservationCapacity entity.ReservationCapacity) error {
	result := d.DB.Create(&reservationCapacity)
	return result.Error
}

func (d *PostgressDataAccess) PostReservationUser(reservationUser entity.ReservationUser) error {
	result := d.DB.Create(&reservationUser)
	return result.Error
}

func (d *PostgressDataAccess) PatchReservationUser(reservationUser entity.ReservationUser) error {
	result := d.DB.Updates(&reservationUser)
	return result.Error
}
