package cron

// AddSecondlyJob adds a schedule job run in every second.
func (c *Cron) AddSecondlyJob(cmd func()) {
	c.core.AddFunc("@every 1s", cmd)
}

// AddMinutelyJob adds a schedule job run in every minute.
func (c *Cron) AddMinutelyJob(cmd func()) {
	c.core.AddFunc("*/1 * * * *", cmd)
}

// AddHourlyJob adds a schedule job run in every hour.
func (c *Cron) AddHourlyJob(cmd func()) {
	c.core.AddFunc("@hourly", cmd)
}

// AddDailyJob adds a schedule job run in every day.
func (c *Cron) AddDailyJob(cmd func()) {
	c.core.AddFunc("@daily", cmd)
}

// AddWeeklyJob adds a schedule job run in every week.
func (c *Cron) AddWeeklyJob(cmd func()) {
	c.core.AddFunc("@weekly", cmd)
}

// AddMonthlyJob adds a schedule job run in every month.
func (c *Cron) AddMonthlyJob(cmd func()) {
	c.core.AddFunc("@monthly", cmd)
}

// AddYearlyJob adds a schedule job run in every year.
func (c *Cron) AddYearlyJob(cmd func()) {
	c.core.AddFunc("@yearly", cmd)
}
