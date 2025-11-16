package scheduler_test

import (
	"testing"
	"time"

	"github.com/alex-cos/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestMinuteNext(t *testing.T) {
	t.Parallel()

	s := scheduler.NewEveryMinute(2, 15)

	now := time.Date(2025, 1, 8, 10, 15, 12, 11, time.UTC)
	next := s.Next(now)

	expected := time.Date(2025, 1, 8, 10, 18, 15, 0, time.UTC)
	assert.Equal(t, expected, next)
}
