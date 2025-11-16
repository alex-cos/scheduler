package scheduler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Cron struct {
	Minutes  map[int]bool
	Hours    map[int]bool
	Days     map[int]bool
	Months   map[int]bool
	Weekdays map[int]bool
}

// NewCron parses a simplified cron expression and returns a Cron.
//
// The expression uses 5 fields, similar to standard cron, but with a reduced syntax:
//
//	MINUTE HOUR DAY MONTH WEEKDAY
//
// Fields:
//   - MINUTE:  0–59
//   - HOUR:    0–23
//   - DAY:     1–31
//   - MONTH:   1–12
//   - WEEKDAY: 0–6 (0 = Sunday)
//
// Supported syntax per field:
//   - matches any value
//     N        a single value (e.g., "3")
//     A,B,C    a list of values (e.g., "1,5,10")
//     A-B      a range of values (e.g., "8-17")
//     */N      a step over the full range (e.g., "*/15" → every 15 units)
//
// Examples:
//
//	"0 3 * * *"        → every day at 03:00
//	"*/5 * * * *"      → every 5 minutes
//	"0 */2 * * *"      → every 2 hours
//	"30 8-18 * * 1-5"  → at xx:30, every hour from 8 to 18, Monday to Friday
//	"0 0 1 * *"        → monthly at midnight on the 1st
//
// No support for seconds, months by name, special cron operators (L, W, #),
// or advanced expressions to keep the syntax small and predictable.
func NewCron(expr string) (*Cron, error) {
	parts := strings.Fields(expr)
	if len(parts) != 5 {
		return nil, errors.New("invalid expression: need 5 fields")
	}

	minutes, err := parseField(parts[0], 0, 59)
	if err != nil {
		return nil, fmt.Errorf("minutes: %w", err)
	}

	hours, err := parseField(parts[1], 0, 23)
	if err != nil {
		return nil, fmt.Errorf("hours: %w", err)
	}

	days, err := parseField(parts[2], 1, 31)
	if err != nil {
		return nil, fmt.Errorf("days: %w", err)
	}

	months, err := parseField(parts[3], 1, 12)
	if err != nil {
		return nil, fmt.Errorf("months: %w", err)
	}

	weekdays, err := parseField(parts[4], 0, 6)
	if err != nil {
		return nil, fmt.Errorf("weekdays: %w", err)
	}

	return &Cron{
		Minutes:  minutes,
		Hours:    hours,
		Days:     days,
		Months:   months,
		Weekdays: weekdays,
	}, nil
}

func (c *Cron) Next(after time.Time) time.Time {
	t := after.Add(time.Minute).Truncate(time.Minute)

	for {
		if c.Months[int(t.Month())] &&
			c.Days[t.Day()] &&
			c.Hours[t.Hour()] &&
			c.Minutes[t.Minute()] &&
			c.Weekdays[int(t.Weekday())] {
			return t
		}

		t = t.Add(time.Minute)
	}
}

// ----------------------------------------------------------------------------
// Unexported functions
// ----------------------------------------------------------------------------

func parseField(field string, minimum, maximum int) (map[int]bool, error) {
	result := make(map[int]bool)

	// Wildcard
	if field == "*" {
		for i := minimum; i <= maximum; i++ {
			result[i] = true
		}
		return result, nil
	}

	// Step (*/N)
	if strings.HasPrefix(field, "*/") {
		step, err := strconv.Atoi(field[2:])
		if err != nil {
			return nil, fmt.Errorf("invalid step: %s", field)
		}
		for i := minimum; i <= maximum; i += step {
			result[i] = true
		}
		return result, nil
	}

	// List or ranges
	parts := strings.Split(field, ",")
	for _, p := range parts {
		// Range (A-B)
		if strings.Contains(p, "-") {
			bounds := strings.SplitN(p, "-", 2)
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			if start > end || start < minimum || end > maximum {
				return nil, fmt.Errorf("invalid range: %s", p)
			}
			for i := start; i <= end; i++ {
				result[i] = true
			}
		} else {
			// Single value
			v, err := strconv.Atoi(p)
			if err != nil || v < minimum || v > maximum {
				return nil, fmt.Errorf("invalid value: %s", p)
			}
			result[v] = true
		}
	}

	return result, nil
}
