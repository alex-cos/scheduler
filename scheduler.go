package scheduler

import (
	"time"
)

type Scheduler struct {
	ticker   *time.Ticker
	schedule Schedule
}

// -----------------------------------------------------------------------------
// Constructor
// -----------------------------------------------------------------------------

func NewScheduler(s Schedule) *Scheduler {
	scheduler := &Scheduler{
		ticker:   nil,
		schedule: s,
	}
	scheduler.ticker = time.NewTicker(scheduler.computeNextTime())

	return scheduler
}

func (s *Scheduler) C() <-chan time.Time {
	return s.ticker.C
}

func (s *Scheduler) Stop() {
	s.ticker.Stop()
}

func (s *Scheduler) Reset() {
	s.ticker.Reset(s.computeNextTime())
}

func (s *Scheduler) GetNextTime() time.Time {
	return s.schedule.Next(time.Now())
}

// ----------------------------------------------------------------------------
// Unexported functions
// ----------------------------------------------------------------------------

func (s *Scheduler) computeNextTime() time.Duration {
	next := s.schedule.Next(time.Now())

	return time.Until(next)
}
