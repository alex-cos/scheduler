# Simple and Predictable Time-Based Scheduling for Go

`github.com/alex-cos/scheduler` is a lightweight, predictable, and extensible
scheduler for Go.

![scheduler](/scheduler.png)

It provides:

- A **mini-cron syntax** (simple, safe, readable)
- Native **Daily**, **Weekly**, and aligned schedules
- A clean `Schedule` interface to build your own timing logic
- A deterministic, minute-resolution scheduler loop

Perfect for periodic jobs, maintenance tasks, data collection, and anything that
requires **clock-based scheduling** without the complexity of full cron.

---

## Installation

```sh
go get github.com/alex-cos/scheduler
```

## Basic Usage

### Run a task every day at 03:00

```go
package main

import (
    "fmt"
    "github.com/alex-cos/scheduler"
)

func main() {
    // every day at 03:00
    cron, err := scheduler.NewCronSchedule("0 3 * * *")
    if (err != nil) {
      panic(err)
    }
    s := scheduler.NewScheduler(cron)

    for {
      select {
      case t := <-s.C():
        fmt.Println("Tick:", t)
        // Put a job here that could take time.
        s.Reset()
      }
    }
}
```

## Mini-Cron Syntax

```txt
MINUTE HOUR DAY MONTH WEEKDAY
```

## Schedules Included

1. CronSchedule

Created using the mini-cron syntax:

```go
cron, _ := scheduler.NewCronSchedule("*/10 * * * *")
```

Triggers every 10 minutes.

2. Daily

Run at a specific hour and minute:

```go
s := scheduler.NewDaily(3, 0, 0) // Every day at 03:00:00
```

3. Weekly

Run once per week at a weekday and time:

```go
s := scheduler.NewWeekly(time.Monday, 9, 30, 0) // Every monday at 09:00:00
```

4. EveryMinute (aligned)

```go
s := scheduler.NewEveryMinute(15, 5)  // Every 15 minutes with a shift of 5 seconds
```

## Writing Your Own Schedule

A schedule only needs to implement:

```go
type Schedule interface {
    Next(after time.Time) time.Time
}
```

Example: run every 90 minutes aligned to the hour:

```go
type Every90Minutes struct{}

func (Every90Minutes) Next(after time.Time) time.Time {
    t := after.Truncate(time.Hour).Add(90 * time.Minute)
    if !t.After(after) {
        t = t.Add(90 * time.Minute)
    }
    return t
}
```

## Tests

Unit tests are provided for:

- cron parsing
- cron schedule resolution
- Weekly schedule
- aligned interval schedules
- field parsing (lists, ranges, steps)

Run tests with:

```sh
go test ./...
```
