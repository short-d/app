package mdtest

import "github.com/short-d/app/fw"

var _ fw.Metrics = (*MetricsFake)(nil)

type MetricsFake struct {
}

func (m MetricsFake) Count(metricID string, point int, interval int, ctx fw.ExecutionContext) {
}

func (m MetricsFake) Rate(metricID string, point float32, interval int, ctx fw.ExecutionContext) {
}

func (m MetricsFake) Gauge(metricID string, point float32, ctx fw.ExecutionContext) {
}

func NewMetricsFake() MetricsFake {
	return MetricsFake{}
}
