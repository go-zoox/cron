package cron

// AddSecondlyJob adds a schedule job run in every second.
func (c *Cron) AddSecondlyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "@every 1s", cmd)
}

// AddMinutelyJob adds a schedule job run in every minute.
func (c *Cron) AddMinutelyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "*/1 * * * *", cmd)
}

// AddHourlyJob adds a schedule job run in every hour.
func (c *Cron) AddHourlyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "@hourly", cmd)
}

// AddDailyJob adds a schedule job run in every day.
func (c *Cron) AddDailyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "@daily", cmd)
}

// AddWeeklyJob adds a schedule job run in every week.
func (c *Cron) AddWeeklyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "@weekly", cmd)
}

// AddMonthlyJob adds a schedule job run in every month.
func (c *Cron) AddMonthlyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "@monthly", cmd)
}

// AddYearlyJob adds a schedule job run in every year.
func (c *Cron) AddYearlyJob(name string, cmd func() error) (id int, err error) {
	return c.AddJob(name, "@yearly", cmd)
}
