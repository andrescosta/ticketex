package entity

import (
	"time"

	"github.com/andrescosta/ticketex/func/internal/ticket/enums"
	"gorm.io/gorm"
)

type TicketTrans struct {
	AdventureID    string `gorm:"primaryKey:Adventure_id"`
	UserID         string `gorm:"primaryKey:User_id"`
	Type           string `gorm:"primaryKey"`
	CreditCardTXID string
	Quantity       uint
	Status         enums.TransactionStatus
	Tickets        []Ticket `gorm:"foreignKey:Adventure_id,User_id,Type"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type Ticket struct {
	AdventureID string `gorm:"primaryKey:Adventure_id"`
	UserID      string `gorm:"primaryKey:User_id"`
	Type        string `gorm:"primaryKey"`
	Code        string `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
