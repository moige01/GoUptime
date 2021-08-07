package main

import (
	"log"
	"os"
)

const (
	DEFAULT_LOG_PATH = "logs/gouptime.log"
)

func setGlobalLogToFile() {
	log_path, ok := os.LookupEnv("LOG_FILE_PATH")

	if !ok {
		log_path = DEFAULT_LOG_PATH
	}

	file, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal("Can't open the desired log file: ", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
