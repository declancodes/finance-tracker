package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// Need to register postgres driver
	_ "github.com/lib/pq"
)

// DbConn is the means of getting a connection to the database.
func DbConn() (db *sqlx.DB) {
	dbUser, dbPassword, dbHost, dbPort, dbName :=
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
