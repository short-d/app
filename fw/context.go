package fw

import "time"

type ExecutionContext struct {
	RequestID        string
	RequestStartAt   time.Time
	FeatureToggleID  string
	ExperimentBucket string
}
