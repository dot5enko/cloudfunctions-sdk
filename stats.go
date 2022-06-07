package sdk

import "time"

type Metric struct {
	Calls uint
	Time  time.Duration
}

func (item *Metric) Call(startTime time.Time) {
	item.Calls += 1
	item.Time += time.Since(startTime)
}
