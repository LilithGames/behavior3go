package actions

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Subscriber struct {
	core.Action
	SubTopic func(tick core.Ticker, client interface{}) error
}

func (s *Subscriber) OnTick(tick core.Ticker) b3.Status {
	value := s.GetValueFromAncestor("subClient", tick.Blackboard())
	if value == nil {
		return b3.FAILURE
	}
	if s.SubTopic == nil {
		return b3.FAILURE
	}
	err := s.SubTopic(tick, value)
	if err != nil {
		return b3.FAILURE
	}
	return b3.SUCCESS
}

func (s *Subscriber) GetClass() string {
	return b3.SUBSCRIBER
}
