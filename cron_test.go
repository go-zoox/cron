package cron

import (
	"sync"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	wg := &sync.WaitGroup{}

	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	// c.AddJob("*/1 * * * *", func() {
	// 	t.Log("cron job ran at", time.Now())
	// })

	wg.Add(1)
	start := time.Now()
	c.AddSecondlyJob("test", func() error {
		t.Log("cron job ran at", time.Now())
		if time.Since(start) > 3*time.Second {
			wg.Done()
		}
		return nil
	})

	c.Start()

	// for {
	// 	time.Sleep(time.Second)
	// }

	wg.Wait()
}
