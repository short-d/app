package metrics

import "github.com/short-d/app/fw"

var _ Metrics = (*Fake)(nil)

type Fake struct {
}

func (f Fake) Count(metricID string, point int, interval int, ctx fw.ExecutionContext) {
}

func (f Fake) Rate(metricID string, point float32, interval int, ctx fw.ExecutionContext) {
}

func (f Fake) Gauge(metricID string, point float32, ctx fw.ExecutionContext) {
}

func NewFake() Fake {
	return Fake{}
}
