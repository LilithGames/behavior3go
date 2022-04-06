package composites

import (
	"context"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Sequence struct {
	core.Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (s *Sequence) OnTick(tick core.Ticker) b3.Status {
	cancelCtx := context.Background()
	if ctx := s.GetValueFromAncestor("cancelCtx", tick.Blackboard()); ctx != nil {
		cancelCtx = ctx.(context.Context)
	}
	for i := 0; i < s.GetChildCount(); i++ {
		var status = s.GetChild(i).Execute(tick)
		for status == b3.RUNNING {
			select {
			case <-cancelCtx.Done():
				status = b3.SUCCESS
			default:
				status = s.GetChild(i).Execute(tick)
			}
		}
		if status != b3.SUCCESS {
			return status
		}
	}
	return b3.SUCCESS
}

func (s *Sequence) GetClass() string {
	return b3.SEQUENCE
}
