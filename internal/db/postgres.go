package db

import (
	"errors"
	"log"
	"mail-service/internal/config"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func PgConnect() error {
	DB, err = gorm.Open(postgres.Open(config.PgConnString()), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return errors.New("failed to initialize connection to postgres db")
	}

	log.Print("Successfully connected to Postgres Database")

	return nil
}
