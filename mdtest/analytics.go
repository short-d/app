package mdtest

import "github.com/short-d/app/fw"

var _ fw.Analytics = (*AnalyticsFake)(nil)

type AnalyticsFake struct {
}

func (a AnalyticsFake) Identify(userID string, traits map[string]string) {
}

func (a AnalyticsFake) Track(eventName string, properties map[string]string, userID string, ctx fw.ExecutionContext) {
}

func (a AnalyticsFake) Group(userID string, groupID string) {
}

func (a AnalyticsFake) Alias(prevUserID string, newUserID string) {
}

func NewAnalyticsFake() AnalyticsFake {
	return AnalyticsFake{}
}
