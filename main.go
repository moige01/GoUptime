// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/go-co-op/gocron"
)

var (
	scheduler *gocron.Scheduler
	interval  string
)

func init() {
	setGlobalLogToFile()
	initLoggers()
	initScheduler(&scheduler, &interval)
}

func main() {
	h := make(Handlers)
	uc := new(UptimeChecker)
	csv := new(CSVDataSource)

	h.RegisterHandler("STDOUT", new(StdoutHandler))
	h.RegisterHandler("LOG", new(LogHandler))
	h.RegisterHandler("DISCORD", new(DiscordHandler))

	_, err := scheduler.SingletonMode().Do(func() {
		uc.Init(csv, h)
		uc.VerifyStatus()
	})

	if err != nil {
		ErrorLogger.Fatal(err)
	}

	scheduler.StartBlocking()
}
