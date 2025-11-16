package scheduler

import "time"

// Weekly triggers the event every week at a specific weekday and time.
//
// Weekday must be in the Go convention:
//   Sunday = 0, Monday = 1, ..., Saturday = 6.
//
// Example:
//   Weekly{Weekday: time.Monday, Hour: 9, Minute: 30}
//   → every Monday at 09:30.
type Weekly struct {
	weekday time.Weekday
	hour    int
	minute  int
	second  int
}

func NewWeekly(weekday time.Weekday, hour, minute, second int) *Weekly {
	return &Weekly{
		weekday: weekday,
		hour:    hour,
		minute:  minute,
		second:  second,
	}
}

func (w Weekly) Next(after time.Time) time.Time {
	// Start from the next minute boundary
	t := after.Add(time.Minute).Truncate(time.Minute)

	// Compute today's scheduled time
	next := time.Date(
		t.Year(), t.Month(), t.Day(),
		w.hour, w.minute, w.second, 0,
		t.Location(),
	)

	// If the scheduled time for this week is already passed → next week
	daysAhead := int(w.weekday - next.Weekday())
	if daysAhead < 0 {
		daysAhead += 7
	}

	next = next.AddDate(0, 0, daysAhead)

	// Check if it's still behind "after"
	if !next.After(after) {
		next = next.AddDate(0, 0, 7)
	}

	return next
}
