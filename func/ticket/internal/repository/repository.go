package repository

import (
	"log"
	"os"
	"time"

	"github.com/andrescosta/ticketex/func/ticket/internal/config"
	"github.com/andrescosta/ticketex/func/ticket/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ITicket interface {
	GetTicketTrans(ticketTransq entity.TicketTrans) (entity.TicketTrans, error)
	NewTicketTrans(ticketTrans entity.TicketTrans) error
	UpdateTicketTrans(ticketTrans entity.TicketTrans) error
}

type Ticket struct {
	DB *gorm.DB
}

func Init(config config.Config) (ITicket, error) {
	dataAccess := &Ticket{}
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
		return nil, err
	}
	dataAccess.DB = db
	return dataAccess, nil
}

func (d *Ticket) GetTicketTrans(ticketTransq entity.TicketTrans) (entity.TicketTrans, error) {
	var ticketTrans entity.TicketTrans
	res := d.DB.Model(&entity.TicketTrans{}).Preload("Tickets").First(&ticketTrans, ticketTransq)
	return ticketTrans, res.Error
}
func (d *Ticket) NewTicketTrans(ticketTrans entity.TicketTrans) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		if res := tx.Create(&ticketTrans); res.Error != nil {
			return res.Error
		}
		return nil
	})
}

func (d *Ticket) UpdateTicketTrans(ticketTrans entity.TicketTrans) error {
	result := d.DB.Updates(&ticketTrans)
	return result.Error
}
