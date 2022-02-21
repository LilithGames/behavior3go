package composites

import (
	"context"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type MemSequence struct {
	core.Composite
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (s *MemSequence) OnOpen(tick core.Ticker) {
	tick.Blackboard().Set("runningChild", 0, tick.GetTree().GetID(), s.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (s *MemSequence) OnTick(tick core.Ticker) b3.Status {
	child := tick.Blackboard().GetInt("runningChild", tick.GetTree().GetID(), s.GetID())
	cancelCtx := context.Background()
	if ctx := s.GetValueFromAncestor("cancelCtx", tick.Blackboard()); ctx != nil {
		cancelCtx = ctx.(context.Context)
	}
	for i := child; i < s.GetChildCount(); i++ {
		tick.Blackboard().Set("runningChild", i, tick.GetTree().GetID(), s.GetID())
		var status = s.GetChild(i).Execute(tick)
		for status == b3.RUNNING {
			select {
			case <- cancelCtx.Done():
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

func (s *MemSequence) GetClass() string {
	return b3.MEMSEQUENCE
}
