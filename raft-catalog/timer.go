package main

import (
	"time"
)

type Timer struct {
	TimerID *time.Timer
	RandTime int
}

func NewTimer() *Timer {
	return &Timer{}
}

// func (t *Timer) Start(randTime int) {
// 	t.RandTime = randTime
// 	t.TimerID = time.NewTimer(time.Duration(randTime) * time.Millisecond)
// }

// func (t *Timer) Stop() {
// 	if t.TimerID != nil {
// 		t.TimerID.Stop()
// 	}
// }

// func (t *Timer) Reset() {
// 	t.Stop()
// 	t.Start(t.RandTime)
// }
