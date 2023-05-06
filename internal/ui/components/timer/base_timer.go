package timer

import (
	"fmt"
	"time"
)

type baseTimer struct {
	startTime time.Time
	paused    bool
	stopped   bool
	total     int64
	elapsed   int64
}

func (timer *baseTimer) Start() {
	timer.startTime = time.Now()
}

func (timer *baseTimer) Pause() error {
	if timer.paused {
		return fmt.Errorf("already paused")
	}
	timer.paused = true
	return nil
}

func (timer *baseTimer) IsPaused() bool {
	return timer.paused
}

func (timer *baseTimer) Stop() error {
	if timer.stopped {
		return fmt.Errorf("already stopped")
	}
	timer.stopped = true
	return nil
}

func (timer *baseTimer) IsStopped() bool {
	return timer.stopped
}

func (timer *baseTimer) update() {
	if timer.paused || timer.stopped {
		return
	}
	currentTime := time.Now()
	timer.elapsed = int64(currentTime.Sub(timer.startTime).Seconds())
}
