package scheduler

import "time"

type EveryHour struct {
	interval     int
	offsetMinute int
}

func NewEveryHour(interval, offsetMinute int) *EveryHour {
	return &EveryHour{
		interval:     interval,
		offsetMinute: offsetMinute,
	}
}

func (e EveryHour) Next(t time.Time) time.Time {
	next := t.Truncate(time.Hour).Add(time.Duration(e.interval) * time.Hour)
	hour := next.Hour()
	mod := hour % e.interval
	if mod != 0 {
		diff := e.interval - mod
		next = next.Add(time.Duration(diff) * time.Hour)
	}
	next = next.Add(time.Duration(e.offsetMinute) * time.Minute)
	if !next.After(t) {
		next = next.Add(time.Duration(e.interval) * time.Hour)
	}

	return next
}
