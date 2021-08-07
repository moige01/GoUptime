package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

const (
	DEFAULT_LOG_PATH = "logs/gouptime.log"
)

func getFlag(cron_interval *string) {
	flag.StringVar(
		cron_interval,
		"interval",
		"",
		"A valid similar UNIX CRON string time interval. Note that this syntax is\n"+
			"enhanced by https://pkg.go.dev/github.com/robfig/cron package.\n"+
			"Please, see their docs to see vailable features."+
			"\n"+
			"Examples:\n"+
			"\t*/5 * * * * * (Every five seconds. Note the extra entry. The first one represent seconds intervals)\n"+
			"\t@every 1h30m (Every one hour thirty. Using time.ParseDuration valid syntax)\n"+
			"\t@daily (Run once a day or midnight)",
	)

	flag.Parse()

	if *cron_interval == "" {
		flag.Usage()

		os.Exit(1)
	}
}

func initScheduler(s **gocron.Scheduler, cron_interval *string) {
	getFlag(cron_interval)

	*s = gocron.NewScheduler(time.UTC)

	(*s).CronWithSeconds(*cron_interval)

	InfoLogger.Printf("Using CRON interval (%s)\n", *cron_interval)
}

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
