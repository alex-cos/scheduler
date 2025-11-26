package scheduler

import "time"

type EveryHour struct {
	interval int
	offset   time.Duration
}

func NewEveryHour(interval int) *EveryHour {
	return &EveryHour{
		interval: interval,
		offset:   0,
	}
}

func NewEveryHourOffsetMinute(interval, offsetMinutes int) *EveryHour {
	return &EveryHour{
		interval: interval,
		offset:   time.Duration(offsetMinutes) * time.Minute,
	}
}

func NewEveryHourOffsetSecond(interval, offsetSeconds int) *EveryHour {
	return &EveryHour{
		interval: interval,
		offset:   time.Duration(offsetSeconds) * time.Second,
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
	next = next.Add(e.offset)
	if !next.After(t) {
		next = next.Add(time.Duration(e.interval) * time.Hour)
	}

	return next
}
