package db

import (
	"caltoph/internal/logger"
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(postgres_uri string) {
	//Get postgres uri from environment variable
	if postgres_uri == "" {
		logger.FatalLogger.Panicln("No postgres uri configured")
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
