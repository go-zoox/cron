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
* [reugn/go-quartz](https://github.com/reugn/go-quartz) - Minimalist and zero-dependency scheduling library for Go
* [go-co-op/gocron](https://github.com/go-co-op/gocron) - Easy and fluent Go cron scheduling
* [robfig/cron](https://github.com/robfig/cron) - a cron library for go

## License
GoZoox is released under the [MIT License](./LICENSE).
