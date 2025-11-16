package scheduler

import "time"

type EveryMinute struct {
	interval     int
	offsetSecond int
}

func NewEveryMinute(interval, offsetSecond int) *EveryMinute {
	return &EveryMinute{
		interval:     interval,
		offsetSecond: offsetSecond,
	}
}

func (e EveryMinute) Next(t time.Time) time.Time {
	next := t.Truncate(time.Minute).Add(time.Duration(e.interval) * time.Minute)
	minute := next.Minute()
	mod := minute % e.interval
	if mod != 0 {
		diff := e.interval - mod
		next = next.Add(time.Duration(diff) * time.Minute)
	}
	next = next.Add(time.Duration(e.offsetSecond) * time.Second)
	if !next.After(t) {
		next = next.Add(time.Duration(e.interval) * time.Minute)
	}

	return next
}
