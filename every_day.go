package scheduler

import "time"

type Daily struct {
	hour   int
	minute int
	second int
}

func NewDaily(hour, minute, second int) *Daily {
	return &Daily{
		hour:   hour,
		minute: minute,
		second: second,
	}
}

func (d Daily) Next(t time.Time) time.Time {
	next := time.Date(
		t.Year(), t.Month(), t.Day(),
		d.hour, d.minute, d.second, 0,
		t.Location(),
	)
	if !next.After(t) {
		next = next.Add(24 * time.Hour)
	}

	return next
}
