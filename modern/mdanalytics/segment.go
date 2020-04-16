package mdanalytics

import (
	"github.com/short-d/app/fw"
	"gopkg.in/segmentio/analytics-go.v3"
)

// Segment API =>
//   https://segment.com/docs/connections/sources/catalog/libraries/server/http-api/

var _ fw.Analytics = (*Segment)(nil)

type Segment struct {
	client analytics.Client
	timer  fw.Timer
	logger fw.Logger
}

func (s Segment) Identify(userID string, traits map[string]string) {
	segmentTraits := analytics.NewTraits()
	for trait, val := range traits {
		segmentTraits.Set(trait, val)
	}

	now := s.timer.Now()
	s.enqueue(analytics.Identify{
		UserId:    userID,
		Traits:    segmentTraits,
		Timestamp: now,
	})
}

func (s Segment) Group(userID string, groupID string) {
	now := s.timer.Now()
	s.enqueue(analytics.Group{
		UserId:    userID,
		GroupId:   groupID,
		Timestamp: now,
	})
}

func (s Segment) Alias(prevUserID string, newUserID string) {
	now := s.timer.Now()
	s.enqueue(analytics.Alias{
		PreviousId: prevUserID,
		UserId:     newUserID,
		Timestamp:  now,
	})
}

func (s Segment) Track(eventName string, properties map[string]string, userID string, ctx fw.ExecutionContext) {
	props := analytics.NewProperties()
	for prop, val := range properties {
		props.Set(prop, val)
	}
	props.Set("feature-toggle", ctx.FeatureToggleID)
	props.Set("experiment-bucket", ctx.ExperimentBucket)

	now := s.timer.Now()
	s.enqueue(analytics.Track{
		Event:      eventName,
		UserId:     userID,
		Properties: props,
		Timestamp:  now,
	})
}

func (s Segment) enqueue(message analytics.Message) {
	err := s.client.Enqueue(message)
	if err == nil {
		return
	}
	s.logger.Error(err)
}

func NewSegment(segmentWriteKey string, timer fw.Timer, logger fw.Logger) Segment {
	client := analytics.New(segmentWriteKey)
	return Segment{
		client: client,
		timer:  timer,
	}
}
