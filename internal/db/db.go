package db

import (
	"caltoph/internal/logger"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	//Get postgres uri from environment variable
	postgres_uri, postgres_uri_present := os.LookupEnv("POSTGRES_URI")
	if !postgres_uri_present {
		logger.FatalLogger.Panicln("POSTGRES_URI not set")
	}

	//Try database connection and ping
	var err error
	db, err = sql.Open("postgres", postgres_uri)
	if err != nil {
		logger.FatalLogger.Println("Connection to database failed")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.FatalLogger.Println("Initial ping to database failed")
		panic(err)
	}
	logger.InfoLogger.Println("Successfully connected to database")
}

// Ping DB. Return true if succeeded
func PingDB() bool {
	if err := db.Ping(); err != nil {
		logger.WarningLogger.Println("db: Could not ping database")
		return false
	}
	logger.DebugLogger.Println("db: Successfully pinged database")
	return true
}
