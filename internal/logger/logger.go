package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var DebugLogger *log.Logger
var InfoLogger *log.Logger
var WarningLogger *log.Logger
var ErrorLogger *log.Logger
var FatalLogger *log.Logger

func Init() {
	//Initialize loggers that are always active
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Init2(loglevel string) {
	if strings.ToLower(loglevel) == "debug" {
		fmt.Println("Using loglevel debug")
		return
	}
	if strings.ToLower(loglevel) == "info" {
		DebugLogger.SetOutput(io.Discard)
		fmt.Println("Using loglevel info")
		return
	}
	if strings.ToLower(loglevel) == "warning" {
		DebugLogger.SetOutput(io.Discard)
		InfoLogger.SetOutput(io.Discard)
		fmt.Println("Using loglevel warning")
		return
	}

	//If loglevel is not empty string now, it's not legal
	if strings.ToLower(loglevel) != "" {
		WarningLogger.Println("logger: Provided loglevel is not a legal loglevel. Using DEBUG instead.")
		return
	}
	fmt.Println("Using default loglevel")
}
