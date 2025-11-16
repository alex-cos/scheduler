package scheduler_test

import (
	"testing"
	"time"

	"github.com/alex-cos/scheduler"
	"github.com/stretchr/testify/assert"
)

func mustParseCron(t *testing.T, expr string) *scheduler.Cron {
	t.Helper()

	c, err := scheduler.NewCron(expr)
	assert.NoError(t, err)

	return c
}

func TestCronDailyAt3AM(t *testing.T) {
	t.Parallel()

	cron := mustParseCron(t, "0 3 * * *")

	now := time.Date(2025, 1, 10, 2, 0, 0, 0, time.UTC)
	next := cron.Next(now)

	expected := time.Date(2025, 1, 10, 3, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}

func TestCronEvery5Minutes(t *testing.T) {
	t.Parallel()

	cron := mustParseCron(t, "*/5 * * * *")

	now := time.Date(2025, 1, 10, 10, 2, 30, 0, time.UTC)
	next := cron.Next(now)

	// Next 5-minute boundary is 10:05
	expected := time.Date(2025, 1, 10, 10, 5, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}

func TestCronWeekdaysAt0830(t *testing.T) {
	t.Parallel()

	cron := mustParseCron(t, "30 8 * * 1-5")

	// Friday at 09:00 → next Monday
	now := time.Date(2025, 1, 10, 9, 0, 0, 0, time.UTC) // Friday
	next := cron.Next(now)

	expected := time.Date(2025, 1, 13, 8, 30, 0, 0, time.UTC) // Monday
	assert.Equal(t, expected, next)
}

func TestCronMonthly(t *testing.T) {
	t.Parallel()

	cron := mustParseCron(t, "0 0 1 * *")

	// Feb 2 → next trigger March 1
	now := time.Date(2025, 2, 2, 12, 0, 0, 0, time.UTC)
	next := cron.Next(now)

	expected := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}
