package config

import (
	"fmt"
	"os"
	"strconv"

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
	fmt.Println(dsn)

	return dsn
}

func ServerPort() int {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		return 5000 // default port
	}

	return port
}
