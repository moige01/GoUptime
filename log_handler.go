// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

type LogHandler struct{}

func (l *LogHandler) getLoggerFor(url string) *log.Logger {
	logger, err := NewFileLogger()

	if err != nil {
		log.Fatal(err)
	}

	schema_regex := regexp.MustCompile(`^(http|s):\/\/`)
	sanitize_url := schema_regex.ReplaceAllString(url, "")

	file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", sanitize_url), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		logger.GetErrorLogger().Printf("Can't create logger for %s due to %s", url, err.Error())

		return nil
	}

	handler := log.New(file, fmt.Sprintf("[%s]: ", url), log.Ldate|log.Ltime|log.Lshortfile)

	return handler
}

func (l *LogHandler) Dispatch(page string, status int, message string, wg *sync.WaitGroup) {
	defer wg.Done()

	pageLogger := l.getLoggerFor(page)

	if pageLogger == nil {
		return
	}

	if !(status >= 200 && status < 300) {
		pageLogger.Printf("%d - %s\n", status, message)
	}
}
