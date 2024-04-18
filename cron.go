package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-zoox/cache"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/safe"
	robCron "github.com/robfig/cron/v3"
)

// Cron is a schedule job, which can be used to run jobs on a schedule.
type Cron struct {
	core  *robCron.Cron
	cache cache.Cache
	sync.Mutex

	cfg *Config
}

// Config is a wrapper of robCron.Config
type Config struct {
	// TimeZone is the timezone in which the cron will run.
	TimeZone string

	// RunNow specifies whether to run the job immediately after adding it.
	RunRightNow bool
}

// New creates a new Cron with the given configuration.
func New(cfg ...*Config) (*Cron, error) {
	var core *robCron.Cron
	if len(cfg) > 1 {
		return nil, fmt.Errorf("cron: only one config is allowed")
	}

	_cfg := &Config{}
	if len(cfg) == 1 && cfg[0] != nil {
		_cfg = cfg[0]
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
		core:  core,
		cache: cache.New(),
		cfg:   _cfg,
	}, nil
}

// AddJobConfig is a configuration for AddJob.
type AddJobConfig struct {
	// RunRightNow specifies whether to run the job immediately after adding it.
	RunRightNow bool
}

// AddJobOption is a function that configures the AddJob.
type AddJobOption func(cfg *AddJobConfig)

// AddJob adds a Job to the Cron to be run on the given schedule.
func (c *Cron) AddJob(id string, spec string, job func() error, opts ...AddJobOption) error {
	c.Lock()
	defer c.Unlock()

	cfg := &AddJobConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	if ok := c.cache.Has(id); ok {
		return fmt.Errorf("cron: job %s already exists", id)
	}

	// run the job immediately
	if c.cfg.RunRightNow {
		go safe.Do(job)
	} else if cfg.RunRightNow {
		go safe.Do(job)
	}

	innerID, err := c.core.AddFunc(spec, func() {
		if err := safe.Do(job); err != nil {
			logger.Error("[cron][name: %s] job failed: %v", id, err)
		}
	})
	if err != nil {
		return err
	}

	if err := c.cache.Set(id, &innerID); err != nil {
		return err
	}

	return nil
}

// RemoveJob removes a Job from the Cron to be run on the given schedule.
func (c *Cron) RemoveJob(id string) (err error) {
	var innerID robCron.EntryID
	if err = c.cache.Get(id, &innerID); err != nil {
		return
	}

	if err = c.cache.Del(id); err != nil {
		return
	}

	c.core.Remove(innerID)
	return
}

// HasJob returns true if the given job exists.
func (c *Cron) HasJob(id string) bool {
	return c.cache.Has(id)
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
