package scheduler

import (
	"sync"
	"time"
)

type Scheduler struct {
	schedule Schedule

	ch chan time.Time

	mu    sync.Mutex
	timer *time.Timer
	stop  chan struct{}
}

// -----------------------------------------------------------------------------
// Constructor
// -----------------------------------------------------------------------------

func NewScheduler(s Schedule) *Scheduler {
	scheduler := &Scheduler{
		schedule: s,
		ch:       make(chan time.Time, 1),
		mu:       sync.Mutex{},
		timer:    nil,
		stop:     make(chan struct{}),
	}

	go scheduler.run()

	return scheduler
}

func (s *Scheduler) C() <-chan time.Time {
	return s.ch
}

func (s *Scheduler) Stop() {
	close(s.stop)
}

func (s *Scheduler) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.timer == nil {
		return
	}

	delay := s.computeDelay()

	if !s.timer.Stop() {
		select {
		case <-s.timer.C:
		default:
		}
	}
	s.timer.Reset(delay)
}

func (s *Scheduler) GetNextTime() time.Time {
	return s.schedule.Next(time.Now())
}

// ----------------------------------------------------------------------------
// Unexported functions
// ----------------------------------------------------------------------------

func (s *Scheduler) computeDelay() time.Duration {
	next := s.schedule.Next(time.Now())
	delay := time.Until(next)
	if delay < 0 {
		delay = 0
	}
	return delay
}

func (s *Scheduler) run() {
	delay := s.computeDelay()

	s.mu.Lock()
	s.timer = time.NewTimer(delay)
	s.mu.Unlock()

	for {
		select {
		case t := <-s.timer.C:
			select {
			case s.ch <- t:
			default: // drop if slow consumer
			}

			delay := s.computeDelay()

			s.mu.Lock()
			if !s.timer.Stop() {
				select {
				case <-s.timer.C:
				default:
				}
			}
			s.timer.Reset(delay)
			s.mu.Unlock()

		case <-s.stop:
			s.mu.Lock()
			if s.timer != nil {
				s.timer.Stop()
			}
			close(s.ch)
			s.mu.Unlock()
			return
		}
	}
}
