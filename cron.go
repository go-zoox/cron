package cron

import (
	"fmt"
	"time"

	"github.com/go-zoox/logger"
	"github.com/go-zoox/safe"
	robCron "github.com/robfig/cron/v3"
)

// Cron is a schedule job, which can be used to run jobs on a schedule.
type Cron struct {
	core *robCron.Cron
}

// Config is a wrapper of robCron.Config
type Config struct {
	TimeZone string
}

// New creates a new Cron with the given configuration.
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

// AddJob adds a Job to the Cron to be run on the given schedule.
func (c *Cron) AddJob(name string, spec string, job func() error) (int, error) {
	id, err := c.core.AddFunc(spec, func() {
		if err := safe.Do(job); err != nil {
			logger.Error("[cron][name: %s] job failed: %v", name, err)
		}
	})
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// RemoveJob removes a Job from the Cron to be run on the given schedule.
func (c *Cron) RemoveJob(id int) (err error) {
	c.core.Remove(robCron.EntryID(id))
	return
}

// ClearJobs clears all jobs.
func (c *Cron) ClearJobs() (err error) {
	for _, entry := range c.core.Entries() {
		c.core.Remove(entry.ID)
	}

	return
}

// Length returns the number of jobs in the given schedule.
func (c *Cron) Length() int {
	return len(c.core.Entries())
}

// Start starts the cron scheduler in a new goroutine.
func (c *Cron) Start() (err error) {
	c.core.Start()
	return
}

// Stop stops the cron scheduler.
func (c *Cron) Stop() (err error) {
	c.core.Stop()
	return
}

// Restart starts the cron scheduler in a new goroutine.
func (c *Cron) Restart() (err error) {
	c.Stop()
	c.Start()
	return
}

// Clear clears the cron scheduler jobs.
func (c *Cron) Clear() (err error) {
	return c.ClearJobs()
}
