package analytics

import (
	"fmt"

	"github.com/short-d/app/fw/ctx"
)

var _ Analytics = (*Fake)(nil)

type Fake struct {
	Analytics []string
}

func (f *Fake) Identify(userID string, traits map[string]string) {
	str := fmt.Sprintf("{message-type: identify, userID: %s, traits: %v}", userID, traits)
	f.Analytics = append(f.Analytics, str)
}

func (f *Fake) Track(eventName string, properties map[string]string, userID string, ctx ctx.ExecutionContext) {
	str := fmt.Sprintf("{message-type: track, userID: %s, event: %s, properties: %v}", userID, eventName, properties)
	f.Analytics = append(f.Analytics, str)
}

func (f *Fake) Group(userID string, groupID string) {
	str := fmt.Sprintf("{message-type: group, userID: %s, groupID: %s}", userID, groupID)
	f.Analytics = append(f.Analytics, str)
}

func (f *Fake) Alias(prevUserID string, newUserID string) {
	str := fmt.Sprintf("{message-type: alias, previousId: %s, userID: %s}", prevUserID, newUserID)
	f.Analytics = append(f.Analytics, str)
}

func NewFake() Fake {
	return Fake{}
}
