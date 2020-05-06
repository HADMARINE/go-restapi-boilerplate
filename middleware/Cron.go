package middleware

import (
	cron "github.com/robfig/cron/v3"
)

// Cron starts cron worker
func Cron() {
	c := cron.New()
	// c.AddFunc("* * * * *", func(){})
	c.Start()
}
