package main

import (
	"log"
	"os"
	"time"

	"github.com/andrescosta/ticketex/func/internal/config"
	"github.com/andrescosta/ticketex/func/internal/reservation/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Println("Starting migration")
	config := config.Load("../../configs/config-r.json")

	loglevel := logger.Error
	if config.DebugSQL {
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
		log.Fatal(err)
	}
	err = db.AutoMigrate(&entity.ReservationMetadata{}, &entity.ReservationCapacity{}, &entity.Reservation{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration ended")
}
