package cron

// AddSecondlyJob adds a schedule job run in every second.
func (c *Cron) AddSecondlyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "@every 1s", cmd)
}

// AddMinutelyJob adds a schedule job run in every minute.
func (c *Cron) AddMinutelyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "*/1 * * * *", cmd)
}

// AddHourlyJob adds a schedule job run in every hour.
func (c *Cron) AddHourlyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "@hourly", cmd)
}

// AddDailyJob adds a schedule job run in every day.
func (c *Cron) AddDailyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "@daily", cmd)
}

// AddWeeklyJob adds a schedule job run in every week.
func (c *Cron) AddWeeklyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "@weekly", cmd)
}

// AddMonthlyJob adds a schedule job run in every month.
func (c *Cron) AddMonthlyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "@monthly", cmd)
}

// AddYearlyJob adds a schedule job run in every year.
func (c *Cron) AddYearlyJob(id string, cmd func() error) (err error) {
	return c.AddJob(id, "@yearly", cmd)
}
