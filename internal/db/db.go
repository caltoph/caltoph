package db

import (
	"caltoph/internal/logger"
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	var err error
	connStr := "postgres://caltoph:9sJV4BhCWQccfKenmrfkMALojsTGY3M3@postgres.checht.de/caltoph?sslmode=verify-full"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.FatalLogger.Println("Connection to database failed")
	}
	logger.InfoLogger.Println("Successfully connected to database")
	err = db.Ping()
	if err != nil {
		logger.FatalLogger.Println("Ping to database failed")
	}
}
