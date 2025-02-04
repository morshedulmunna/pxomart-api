package utils

import (
	"log"
	"os"
	"time"
)

func SetupLogging() *log.Logger {
	// Create 'logs' directory if it doesn't exist
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatal("Error creating logs directory: ", err)
	}

	// Generate the log file name based on the current date
	logFileName := "logs/" + time.Now().Format("2006-01-02") + ".log"

	// Create/Open the log file in append mode
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	// Set up the logger to write to the log file
	logger := log.New(file, "", log.LstdFlags)
	log.SetOutput(file) // Log to the file by default

	return logger
}
