package entity

import (
	"time"

	"github.com/andrescosta/ticketex/func/reservation/internal/enums"
	"gorm.io/gorm"
)

type ReservationMetadata struct {
	Adventure_id string `gorm:"primaryKey"`
	Status       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt        `gorm:"index"`
	Capacities   []ReservationCapacity `gorm:"foreignKey:Adventure_id"`
}

type ReservationCapacity struct {
	Adventure_id string `gorm:"primaryKey"`
	Type         string `gorm:"primaryKey"`
	Status       enums.ReservationMetadataStatus
	Availability uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Reservation struct {
	Adventure_id string `gorm:"primaryKey"`
	User_id      string `gorm:"primaryKey"`
	Type         string `gorm:"primaryKey"`
	Quantity     uint
	Status       enums.ReservationStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
