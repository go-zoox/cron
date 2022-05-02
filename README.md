# Cron - Make Job Scheduling Easier

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/cron)](https://pkg.go.dev/github.com/go-zoox/cron)
[![Build Status](https://github.com/go-zoox/cron/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/cron/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/cron)](https://goreportcard.com/report/github.com/go-zoox/cron)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/cron/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/cron?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/cron.svg)](https://github.com/go-zoox/cron/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/cron.svg?label=Release)](https://github.com/go-zoox/cron/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/cron
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/cron"
)

func main(t *testing.T) {
	c, err := cron.New()
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
```

## Inspired by
* [dirkaholic/kyoo](https://github.com/dirkaholic/kyoo) - Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries

## Related
* [go-zoox/cocurrent](https://github.com/go-zoox/cocurrent) - A Simple Goroutine Limit Pool
* [go-zoox/waitgroup](https://github.com/go-zoox/waitgroup) - Parallel-Controlled WaitGroup
* [go-zoox/promise](https://github.com/go-zoox/promise) - JavaScript Promise Like with Goroutines

## License
GoZoox is released under the [MIT License](./LICENSE).
