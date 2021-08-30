package composites

import (
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
func (s *MemSequence) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), s.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (s *MemSequence) OnTick(tick *core.Tick) b3.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), s.GetID())
	for i := child; i < s.GetChildCount(); i++ {
		var status = s.GetChild(i).Execute(tick)

		if status != b3.SUCCESS {
			if status == b3.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), s.GetID())
			}

			return status
		}
	}
	return b3.SUCCESS
}

func (s *MemSequence) GetClass() string {
	return b3.MEMSEQUENCE
}
