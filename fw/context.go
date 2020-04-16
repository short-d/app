package fw

import "time"

type ExecutionContext struct {
	RequestID        string
	RequestStartAt   time.Time
	Location         Location
	FeatureToggleID  string
	ExperimentBucket string
}
