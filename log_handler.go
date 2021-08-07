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

var logger *log.Logger

func (l *LogHandler) sanitizeUrl(url string) string {
	schema_regex := regexp.MustCompile(`^(http|s):\/\/`)
	sanitize_url := schema_regex.ReplaceAllString(url, "")

	return sanitize_url
}

func (l *LogHandler) setPrefixForPage(url string) {
	logger.SetPrefix(fmt.Sprintf("[%s]: ", url))
}

func (l *LogHandler) setFileForPage(url string) {
	sanitize_url := l.sanitizeUrl(url)

	file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", sanitize_url), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		ErrorLogger.Printf("Can't create logger for %s due to %s", url, err.Error())

		return
	}

	logger.SetOutput(file)
}

func (l *LogHandler) initLoggerForPage(url string) {
	l.setFileForPage(url)
	l.setPrefixForPage(url)
}

func (l *LogHandler) Dispatch(page string, status int, message string, wg *sync.WaitGroup) {
	defer wg.Done()

	if !(status >= 200 && status < 300) {
		l.initLoggerForPage(page)

		logger.Printf("%d - %s\n", status, message)
	}
}

func init() {
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}
