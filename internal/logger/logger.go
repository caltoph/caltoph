package logger

import (
	"io"
	"log"
	"os"
)

var DebugLogger *log.Logger
var InfoLogger *log.Logger
var WarningLogger *log.Logger
var ErrorLogger *log.Logger
var FatalLogger *log.Logger

func Init() {
	var initialized bool
	var logLevel string
	//Get logging variable or assume DEBUG if not set
	_, logLevelPresent := os.LookupEnv("LOGLEVEL")
	if !logLevelPresent {
		logLevel = "DEBUG"
	}
	if logLevelPresent {
		logLevel = "DEBUG"
	}
	//Initialize loggers that are always active
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	//Initialize other loggers according to loglevel
	if logLevel == "INFO" {
		DebugLogger = log.New(io.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		initialized = true
	}
	if logLevel == "DEBUG" {
		DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		initialized = true
	}
	if !initialized {
		DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		initialized = true
		logLevel = "DEBUG"
	}
	InfoLogger.Println("logger: Initialized with loglevel " + logLevel)
}
