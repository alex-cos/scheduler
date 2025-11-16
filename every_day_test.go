package scheduler_test

import (
	"testing"
	"time"

	"github.com/alex-cos/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestEveryDayNextEqual(t *testing.T) {
	t.Parallel()

	s := scheduler.NewDaily(2, 15, 0)

	now := time.Date(2025, 1, 8, 2, 15, 0, 0, time.UTC)
	next := s.Next(now)

	expected := time.Date(2025, 1, 9, 2, 15, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}

func TestEveryDayNextAfter(t *testing.T) {
	t.Parallel()

	s := scheduler.NewDaily(2, 15, 0)

	now := time.Date(2025, 1, 8, 4, 20, 12, 11, time.UTC)
	next := s.Next(now)

	expected := time.Date(2025, 1, 9, 2, 15, 0, 0, time.UTC)
	assert.Equal(t, expected, next)
}
