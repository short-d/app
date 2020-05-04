package metrics

import (
	"github.com/short-d/app/fw/ctx"
)

// TODO(issue#85): Fill in fake metrics to facilitate testing
var _ Metrics = (*Fake)(nil)

type Fake struct {
}

func (f Fake) Count(metricID string, point int, interval int, ctx ctx.ExecutionContext) {
}

func (f Fake) Rate(metricID string, point float32, interval int, ctx ctx.ExecutionContext) {
}

func (f Fake) Gauge(metricID string, point float32, ctx ctx.ExecutionContext) {
}

func NewFake() Fake {
	return Fake{}
}
