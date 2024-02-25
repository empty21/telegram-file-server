package cache

import (
	"github.com/go-co-op/gocron/v2"
	"telegram-file-server/pkg/log"
	"time"
)

func init() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Panic(err)
	}
	_, err = s.NewJob(
		gocron.DurationJob(time.Hour),
		gocron.NewTask(clearCache),
	)
	if err != nil {
		log.Panic(err)
	}
	log.Info("Cron job started")
	s.Start()
}
