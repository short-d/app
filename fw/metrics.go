package fw

type Metrics interface {
	// Count calculates the sum of all points
	Count(metricID string, point int, interval int)
	// Rate calculates the average of all points
	Rate(metricID string, point float32, interval int)
	// Gauge takes the most recent point
	Gauge(metricID string, point float32)
}
