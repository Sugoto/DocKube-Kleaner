package logger

import (
	"log"
	"os"
)

var (
	LogFile  *os.File
	Logger   *log.Logger
)

func InitLogger() {
	var err error
	LogFile, err = os.OpenFile("cleanup.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(LogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(Logger.Writer())
}
