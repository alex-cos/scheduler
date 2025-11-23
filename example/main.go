package main

import (
	"fmt"

	"github.com/alex-cos/scheduler"
)

func main() {
	// every minute with five seconds offset
	cron := scheduler.NewEveryMinute(1, 5)
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
