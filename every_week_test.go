package scheduler_test

import (
	"testing"
	"time"

	"github.com/alex-cos/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestWeeklyNextMonday0930(t *testing.T) {
	t.Parallel()

	s := scheduler.NewWeekly(time.Monday, 9, 30, 0)

	// We are Wednesday 10:00 → next Monday 09:30
	now := time.Date(2025, 1, 8, 10, 0, 0, 0, time.UTC) // Wed
	next := s.Next(now)

	expected := time.Date(2025, 1, 13, 9, 30, 0, 0, time.UTC) // Monday
	assert.Equal(t, expected, next)
}

func TestWeeklySameDayLater(t *testing.T) {
	t.Parallel()

	s := scheduler.NewWeekly(time.Wednesday, 15, 0, 0)

	// Wednesday at 10:00 → later today at 15:00
	now := time.Date(2025, 1, 8, 10, 0, 0, 0, time.UTC) // Wed
	next := s.Next(now)

	expected := time.Date(2025, 1, 8, 15, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}

func TestWeeklySameDayAlreadyPassed(t *testing.T) {
	t.Parallel()

	s := scheduler.NewWeekly(time.Wednesday, 9, 0, 0)

	// Wednesday 10:00 → next week Wednesday
	now := time.Date(2025, 1, 8, 10, 0, 0, 0, time.UTC)
	next := s.Next(now)

	expected := time.Date(2025, 1, 15, 9, 0, 0, 0, time.UTC) // +7 days
	assert.Equal(t, expected, next)
}
