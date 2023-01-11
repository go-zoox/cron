package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-zoox/cron"
)

func main() {
	wg := sync.WaitGroup{}
	c, err := cron.New()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	wg.Add(1)

	c.AddSecondlyJob("every 1 seconds run once", func() error {
		fmt.Println("hello zero")
		return nil
	})

	fmt.Println("start cron")
	c.Start()

	time.Sleep(3 * time.Second)
	fmt.Println("stop cron")
	c.Stop()

	time.Sleep(3 * time.Second)
	fmt.Println("restart cron")
	c.Restart()

	time.Sleep(3 * time.Second)
	fmt.Println("clear cron")
	c.Clear()

	time.Sleep(3 * time.Second)
	fmt.Println("restart cron")
	c.Restart()

	wg.Wait()
}
