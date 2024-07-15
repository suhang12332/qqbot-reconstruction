package util

import (
	"github.com/roylee0704/gron"
	"time"
)

var c *gron.Cron

func init() {
	c = gron.New()
}

func Subscribe(duration time.Duration, specTime string, fn func()) {
	c.AddFunc(gron.Every(duration).At(specTime), fn)
	c.Start()
}
