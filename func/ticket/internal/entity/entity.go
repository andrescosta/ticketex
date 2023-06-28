package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	Adventure_id string `gorm:"primaryKey"`
	Status       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type ReservationCapacity struct {
	Adventure_id string `gorm:"primaryKey"`
	Type         string `gorm:"primaryKey"`
	Status       uint
	Availability uint
	Max          uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type ReservationUser struct {
	Adventure_id string `gorm:"primaryKey"`
	User_id      string `gorm:"primaryKey"`
	Type         string `gorm:"primaryKey"`
	Quantity     uint
	Status       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
