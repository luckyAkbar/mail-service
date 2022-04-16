package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func PgConnString() string {
	host := os.Getenv("PGHOST")
	db := os.Getenv("PGDATABASE")
	user := os.Getenv("PGUSER")
	pw := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")

	if os.Getenv("ENV") == "production" {
		host = "host.docker.internal" // in the production, will be using docker.
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pw, db, port)

	return dsn
}

func ServerPort() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		return "5000" // default port
	}

	return port
}

func LogLevel() string {
	level := os.Getenv("LOG_LEVEL")

	if level == "" {
		return "debug"
	}

	return level
}
