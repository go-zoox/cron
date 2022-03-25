package cron

import (
	"fmt"
	"time"

	robCron "github.com/robfig/cron/v3"
)

type Cron struct {
	core *robCron.Cron
}

type Config struct {
	TimeZone string
}

func New(cfg ...*Config) (*Cron, error) {
	var core *robCron.Cron
	if len(cfg) > 1 {
		return nil, fmt.Errorf("cron: only one config is allowed")
	}

	if len(cfg) == 1 && cfg[0] != nil {
		_cfg := *cfg[0]
		if _cfg.TimeZone != "" {
			tz, err := time.LoadLocation(_cfg.TimeZone)
			if err != nil {
				return nil, fmt.Errorf("failed to load timezone: %w", err)
			}

			core = robCron.New(robCron.WithLocation(tz))
		}
	} else {
		core = robCron.New()
	}

	return &Cron{
		core: core,
	}, nil
}

func (c *Cron) AddJob(spec string, job func()) {
	c.core.AddFunc(spec, job)
}

func (c *Cron) Start() {
	c.core.Start()
}

func (c *Cron) Stop() {
	c.core.Stop()
}

func (c *Cron) AddSecondlyJob(cmd func()) {
	c.core.AddFunc("@every 1s", cmd)
}

func (c *Cron) AddMinutelyJob(cmd func()) {
	c.core.AddFunc("*/1 * * * *", cmd)
}

func (c *Cron) AddHourlyJob(cmd func()) {
	c.core.AddFunc("@hourly", cmd)
}

func (c *Cron) AddDailyJob(cmd func()) {
	c.core.AddFunc("@daily", cmd)
}

func (c *Cron) AddWeeklyJob(cmd func()) {
	c.core.AddFunc("@weekly", cmd)
}

func (c *Cron) AddMonthlyJob(cmd func()) {
	c.core.AddFunc("@monthly", cmd)
}

func (c *Cron) AddYearlyJob(cmd func()) {
	c.core.AddFunc("@yearly", cmd)
}
