// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type FileLogger struct {
	file *os.File
}

func NewFileLogger() (*FileLogger, error) {
	file, err := os.OpenFile("logs/gouptime.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return nil, err
	}

	return &FileLogger{file: file}, nil
}

func (l *FileLogger) GetInfoLogger() *log.Logger {
	return l.GetLogger("INFO")
}

func (l *FileLogger) GetWarningLogger() *log.Logger {
	return l.GetLogger("WARNING")
}

func (l *FileLogger) GetErrorLogger() *log.Logger {
	return l.GetLogger("ERROR")
}

func (l *FileLogger) GetDebugLogger() *log.Logger {
	return l.GetLogger("DEBUG")
}

func (l *FileLogger) GetLogger(logType string) *log.Logger {
	format := fmt.Sprintf("%s: ", logType)

	return log.New(l.file, format, log.Ldate|log.Ltime|log.Lshortfile)
}
