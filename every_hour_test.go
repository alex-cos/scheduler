package scheduler_test

import (
	"testing"
	"time"

	"github.com/alex-cos/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestHourNextEqual(t *testing.T) {
	t.Parallel()

	s := scheduler.NewEveryHour(2)

	now := time.Date(2025, 1, 8, 2, 12, 0, 0, time.UTC)
	next := s.Next(now)

	expected := time.Date(2025, 1, 8, 4, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, next)

	s = scheduler.NewEveryHourOffsetMinute(1, 12)

	now = time.Date(2025, 1, 8, 2, 12, 0, 0, time.UTC)
	next = s.Next(now)

	expected = time.Date(2025, 1, 8, 3, 12, 0, 0, time.UTC)
	assert.Equal(t, expected, next)

	s = scheduler.NewEveryHourOffsetSecond(2, 30)

	now = time.Date(2025, 1, 8, 2, 12, 0, 0, time.UTC)
	next = s.Next(now)

	expected = time.Date(2025, 1, 8, 4, 0, 30, 0, time.UTC)
	assert.Equal(t, expected, next)
}

func TestEverHourNextAfter(t *testing.T) {
	t.Parallel()

	s := scheduler.NewEveryHourOffsetMinute(4, 12)

	now := time.Date(2025, 1, 8, 2, 32, 0, 0, time.UTC)
	next := s.Next(now)

	expected := time.Date(2025, 1, 8, 8, 12, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}
