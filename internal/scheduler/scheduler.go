package scheduler

import (
	"time"
)

func ScheduleCleanup(interval time.Duration, cleanupFunc func()) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cleanupFunc()
	}
}
