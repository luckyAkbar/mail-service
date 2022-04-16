package db

import (
	"errors"
	"mail-service/internal/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func PgConnect() error {
	DB, err = gorm.Open(postgres.Open(config.PgConnString()), &gorm.Config{})

	if err != nil {
		logrus.Error(err.Error())
		return errors.New("failed to initialize connection to postgres db")
	}

	logrus.Info("Successfully connected to Postgres Database")

	return nil
}
