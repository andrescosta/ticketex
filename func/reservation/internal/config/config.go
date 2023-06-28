package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	PostgressDsn string `json:"postgress_dsn"`
	Host         string `json:"host"`
	DebugSql     bool   `json:"debug_sql"`
}

func Load(fileName string) Config {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open config file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to decode config file:", err)
	}

	return config
}
