package metrics

import "github.com/short-d/app/fw"

type Metrics interface {
	// Count calculates the sum of all points
	Count(metricID string, point int, interval int, ctx fw.ExecutionContext)
	// Rate calculates the average of all points
	Rate(metricID string, point float32, interval int, ctx fw.ExecutionContext)
	// Gauge takes the most recent point
	Gauge(metricID string, point float32, ctx fw.ExecutionContext)
}
