package cron

// AddSecondlyJob adds a schedule job run in every second.
func (c *Cron) AddSecondlyJob(cmd func() error) {
	c.AddJob("@every 1s", cmd)
}

// AddMinutelyJob adds a schedule job run in every minute.
func (c *Cron) AddMinutelyJob(cmd func() error) {
	c.AddJob("*/1 * * * *", cmd)
}

// AddHourlyJob adds a schedule job run in every hour.
func (c *Cron) AddHourlyJob(cmd func() error) {
	c.AddJob("@hourly", cmd)
}

// AddDailyJob adds a schedule job run in every day.
func (c *Cron) AddDailyJob(cmd func() error) {
	c.AddJob("@daily", cmd)
}

// AddWeeklyJob adds a schedule job run in every week.
func (c *Cron) AddWeeklyJob(cmd func() error) {
	c.AddJob("@weekly", cmd)
}

// AddMonthlyJob adds a schedule job run in every month.
func (c *Cron) AddMonthlyJob(cmd func() error) {
	c.AddJob("@monthly", cmd)
}

// AddYearlyJob adds a schedule job run in every year.
func (c *Cron) AddYearlyJob(cmd func() error) {
	c.AddJob("@yearly", cmd)
}
