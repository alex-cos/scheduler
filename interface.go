package scheduler

import "time"

type Schedule interface {
	Next(after time.Time) time.Time
}
