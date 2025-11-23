package scheduler_test

import (
	"testing"
	"time"

	"github.com/alex-cos/scheduler"
)

// ----------------------------------------------------
// Fake schedules
// ----------------------------------------------------

type scheduleFixed struct {
	delay time.Duration
}

func (s scheduleFixed) Next(after time.Time) time.Time {
	return after.Add(s.delay)
}

// ----------------------------------------------------
// Tests
// ----------------------------------------------------

func TestSchedulerFirstTick(t *testing.T) {
	t.Parallel()

	s := scheduler.NewScheduler(scheduleFixed{delay: 20 * time.Millisecond})
	defer s.Stop()

	select {
	case tm := <-s.C():
		if tm.IsZero() {
			t.Fatalf("expected non-zero tick time")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("scheduler did not emit first tick in time")
	}
}

func TestSchedulerMultipleTicks(t *testing.T) {
	t.Parallel()

	s := scheduler.NewScheduler(scheduleFixed{delay: 10 * time.Millisecond})
	defer s.Stop()

	count := 0

Loop:
	for {
		select {
		case <-s.C():
			count++
			if count == 3 {
				break Loop
			}
			s.Reset()
		case <-time.After(300 * time.Millisecond):
			t.Fatalf("expected 3 ticks, got %d", count)
		}
	}
}

// ----------------------------------------------------
// Tests Stop()
// ----------------------------------------------------

func TestSchedulerStopClosesChannel(t *testing.T) {
	t.Parallel()

	s := scheduler.NewScheduler(scheduleFixed{delay: 5 * time.Millisecond})

	select {
	case <-s.C():
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("scheduler did not emit tick")
	}

	s.Stop()

	select {
	case _, ok := <-s.C():
		if ok {
			t.Fatalf("expected channel to be closed")
		}
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("C() should be closed after Stop()")
	}
}
