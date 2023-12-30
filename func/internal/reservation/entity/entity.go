package entity

import (
	"time"

	"github.com/andrescosta/ticketex/func/internal/reservation/enum"
	"gorm.io/gorm"
)

type ReservationMetadata struct {
	AdventureID string `gorm:"primaryKey"`
	Status      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt        `gorm:"index"`
	Capacities  []ReservationCapacity `gorm:"foreignKey:Adventure_id"`
}

type ReservationCapacity struct {
	AdventureID  string `gorm:"primaryKey"`
	Type         string `gorm:"primaryKey"`
	Status       enum.ReservationMetadataStatus
	Availability uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Reservation struct {
	AdventureID string `gorm:"primaryKey"`
	UserID      string `gorm:"primaryKey"`
	Type        string `gorm:"primaryKey"`
	Quantity    uint
	Status      enum.ReservationStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
