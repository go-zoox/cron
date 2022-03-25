package cron

import (
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	// c.AddJob("*/1 * * * *", func() {
	// 	t.Log("cron job ran at", time.Now())
	// })

	c.AddSecondlyJob(func() {
		t.Log("cron job ran at", time.Now())
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
