package ctx

import (
	"time"

	"github.com/short-d/app/fw/geo"
)

type ExecutionContext struct {
	RequestID        string
	RequestStartAt   time.Time
	Location         geo.Location
	FeatureToggleID  string
	ExperimentBucket string
}
