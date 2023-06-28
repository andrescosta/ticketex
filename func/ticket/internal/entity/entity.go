package entity

import (
	"time"

	"github.com/andrescosta/ticketex/func/ticket/internal/enums"
	"gorm.io/gorm"
)

type TicketTrans struct {
	Adventure_id      string `gorm:"primaryKey"`
	User_id           string `gorm:"primaryKey"`
	Type              string `gorm:"primaryKey"`
	Credit_Card_TX_ID string
	Quantity          uint
	Status            enums.TransactionStatus
	Tickets           []Ticket `gorm:"foreignKey:Adventure_id,User_id,Type"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Ticket struct {
	Adventure_id string `gorm:"primaryKey"`
	User_id      string `gorm:"primaryKey"`
	Type         string `gorm:"primaryKey"`
	Code         string `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
