package cron

import (
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	wg := &sync.WaitGroup{}

	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)
	start := time.Now()
	err = c.AddSecondlyJob("test", func() error {
		t.Log("cron job ran at", time.Now())
		if time.Since(start) > 3*time.Second {
			wg.Done()
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	err = c.AddSecondlyJob("test", func() error {
		t.Log("cron job ran at", time.Now())
		if time.Since(start) > 3*time.Second {
			wg.Done()
		}
		return nil
	})
	if err == nil {
		t.Fatal("expected error")
	}

	c.Start()

	wg.Wait()
}

func TestRemoveJobRemovesFromCore(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	if err := c.AddJob("myjob", "@every 1h", func() error {
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if !c.HasJob("myjob") {
		t.Fatal("expected job to exist after add")
	}
	if c.Length() != 1 {
		t.Fatalf("expected 1 entry in core, got %d", c.Length())
	}

	if err := c.RemoveJob("myjob"); err != nil {
		t.Fatal(err)
	}

	if c.HasJob("myjob") {
		t.Fatal("expected job to be removed from cache")
	}
	if c.Length() != 0 {
		t.Fatalf("expected 0 entries in core after remove, got %d", c.Length())
	}
}

func TestRemoveJobReAddAfterRemove(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	var count1, count2 atomic.Int32

	if err := c.AddJob("myjob", "@every 1h", func() error {
		count1.Add(1)
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if err := c.RemoveJob("myjob"); err != nil {
		t.Fatal(err)
	}

	if c.HasJob("myjob") {
		t.Fatal("expected job to be removed from cache")
	}
	if c.Length() != 0 {
		t.Fatalf("expected 0 entries in core after remove, got %d", c.Length())
	}

	if err := c.AddJob("myjob", "@every 1h", func() error {
		count2.Add(1)
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if !c.HasJob("myjob") {
		t.Fatal("expected job to exist after re-add")
	}
	if c.Length() != 1 {
		t.Fatalf("expected exactly 1 entry in core after re-add, got %d", c.Length())
	}
}

func TestRemoveNonexistentJobNoop(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	if err := c.RemoveJob("nonexistent"); err != nil {
		t.Fatal("RemoveJob should not error on nonexistent job")
	}

	if c.Length() != 0 {
		t.Fatalf("expected 0 entries, got %d", c.Length())
	}
}

func TestNewWithExplicitTimeZone(t *testing.T) {
	c, err := New(&Config{TimeZone: "Asia/Shanghai"})
	if err != nil {
		t.Fatal(err)
	}
	if c.core == nil {
		t.Fatal("core should not be nil")
	}
	loc := c.core.Location()
	if loc.String() != "Asia/Shanghai" {
		t.Fatalf("expected Asia/Shanghai, got %s", loc.String())
	}
}

func TestNewFallbackToTZEnv(t *testing.T) {
	os.Setenv("TZ", "Asia/Shanghai")
	defer os.Unsetenv("TZ")

	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if c.core == nil {
		t.Fatal("core should not be nil")
	}
	loc := c.core.Location()
	if loc.String() != "Asia/Shanghai" {
		t.Fatalf("expected Asia/Shanghai from TZ env, got %s", loc.String())
	}
}

func TestNewExplicitTimeZoneOverridesTZEnv(t *testing.T) {
	os.Setenv("TZ", "Asia/Shanghai")
	defer os.Unsetenv("TZ")

	c, err := New(&Config{TimeZone: "UTC"})
	if err != nil {
		t.Fatal(err)
	}
	loc := c.core.Location()
	if loc.String() != "UTC" {
		t.Fatalf("expected UTC from explicit config over TZ env, got %s", loc.String())
	}
}

func TestNewNoTimeZoneNoEnv(t *testing.T) {
	os.Unsetenv("TZ")

	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if c.core == nil {
		t.Fatal("core should not be nil")
	}
}

func TestNewInvalidTimeZonePanics(t *testing.T) {
	_, err := New(&Config{TimeZone: "Invalid/TimeZone"})
	if err == nil {
		t.Fatal("expected error for invalid timezone")
	}
}

func TestNewInvalidTZEnvPanics(t *testing.T) {
	os.Setenv("TZ", "Invalid/TimeZone")
	defer os.Unsetenv("TZ")

	_, err := New()
	if err == nil {
		t.Fatal("expected error for invalid TZ env")
	}
}

func TestAddDuplicateJobError(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	if err := c.AddJob("dup", "@every 1h", func() error { return nil }); err != nil {
		t.Fatal(err)
	}
	if err := c.AddJob("dup", "@every 1h", func() error { return nil }); err == nil {
		t.Fatal("expected error for duplicate job")
	}
}

func TestHasJob(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	if c.HasJob("any") {
		t.Fatal("expected false for nonexistent job")
	}

	if err := c.AddJob("any", "@every 1h", func() error { return nil }); err != nil {
		t.Fatal(err)
	}
	if !c.HasJob("any") {
		t.Fatal("expected true after add")
	}

	if err := c.RemoveJob("any"); err != nil {
		t.Fatal(err)
	}
	if c.HasJob("any") {
		t.Fatal("expected false after remove")
	}
}

func TestConcurrentAddRemove(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	const numWorkers = 10
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_ = c.AddJob("job", "@every 1h", func() error { return nil })
				_ = c.RemoveJob("job")
			}
		}(i)
	}

	wg.Wait()
}
