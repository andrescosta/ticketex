package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type Config struct {
	PostgressDsn string `json:"postgress_dsn"`
	Host         string `json:"host"`
	DebugSQL     bool   `json:"debug_sql"`
}

func Load(fileName string) Config {
	fileName = filepath.Clean(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Print("Failed to open config file.")
		config := Config{}
		config = overwriteConfigUsingEnv(config)
		return config
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing config file %v", err)
		}
	}()
	config, err := decode(file)
	if err != nil {
		log.Print("Failed to parse config file:", err)
		config = Config{}
	}
	config = overwriteConfigUsingEnv(config)
	return config
}

func decode(file *os.File) (Config, error) {
	decoder := json.NewDecoder(file)
	var config Config
	err := decoder.Decode(&config)
	return config, err
}

func overwriteConfigUsingEnv(config Config) Config {
	newconfig := config
	v := reflect.ValueOf(config)
	t := v.Type()
	elem := reflect.ValueOf(&newconfig).Elem()
	for i := 0; i < v.NumField(); i++ {
		p := t.Field(i).Tag.Get("json")
		if value := os.Getenv(p); value != "" {
			name := t.Field(i).Name
			ee := elem.FieldByName(name)
			if ee.Kind() == reflect.String {
				ee.SetString(value)
			} else if ee.Kind() == reflect.Bool {
				if bool1, err := strconv.ParseBool(value); err == nil {
					ee.SetBool(bool1)
				}
			} else if ee.Kind() == reflect.Int {
				if int1, err := strconv.Atoi(value); err != nil {
					ee.SetInt(int64(int1))
				}
			}
		}
	}
	return newconfig
}
