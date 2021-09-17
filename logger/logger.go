package logger

import (
	"log"
	"os"
)

var (
	// WarningLogger log warnings to standard output
	WarningLogger *log.Logger
	// InfoLogger log generic information to standard output
	InfoLogger *log.Logger
	// ErrorLogger log errors to standard error output
	ErrorLogger *log.Logger
)

// InitializeLogger method to initialize custom logger for `ping` service
func InitializeLogger() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}
