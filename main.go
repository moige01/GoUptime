// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	h := make(Handlers)
	uc := new(UptimeChecker)

	h.registerHandler("STDOUT", new(StdoutHandler))
	h.registerHandler("LOG", new(LogHandler))

	s.Every(5).Minutes().Do(func() {
		uc.Init(new(CSVDataPopulation), h)
		uc.VerifyStatus()
	})

	s.StartBlocking()
}
