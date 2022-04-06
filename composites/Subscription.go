package composites

import (
	"context"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type SubClient interface {
	Run() error
	Close() error
}

type Subscription struct {
	core.Composite
	ClientCreator func(tick core.Ticker) SubClient
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (s *Subscription) OnTick(tick core.Ticker) b3.Status {
	if !s.matchCondition() {
		return b3.FAILURE
	}
	if s.ClientCreator == nil {
		return b3.FAILURE
	}
	client := s.ClientCreator(tick)
	tick.Blackboard().Set("subClient", client, s.GetTreeID(), s.GetID())
	for i := 0; i < s.GetChildCount(); i++ {
		var status = s.GetChild(i).Execute(tick)
		if status != b3.SUCCESS {
			return status
		}
	}
	ctxValue := s.GetValueFromAncestor("cancelCtx", tick.Blackboard())
	if ctxValue == nil {
		return b3.FAILURE
	}
	ctx := ctxValue.(context.Context)
	go client.Run()
	<-ctx.Done()
	client.Close()
	return b3.SUCCESS
}

func (s *Subscription) GetClass() string {
	return b3.SUBSCRIPTION
}

func (s *Subscription) matchCondition() bool {
	parent := s.GetParent()
	for {
		if parent == nil {
			return false
		}
		if parent.GetClass() == b3.PARALLEL {
			return true
		}
		if parent.GetClass() == b3.SUBSCRIPTION {
			return false
		}
		parent = parent.GetParent()
	}
}
