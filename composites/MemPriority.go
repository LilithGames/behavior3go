package composites

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type MemPriority struct {
	core.Composite
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (p *MemPriority) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), p.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (p *MemPriority) OnTick(tick *core.Tick) b3.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), p.GetID())
	for i := child; i < p.GetChildCount(); i++ {
		var status = p.GetChild(i).Execute(tick)

		if status != b3.FAILURE {
			if status == b3.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), p.GetID())
			}

			return status
		}
	}
	return b3.FAILURE
}

func (p *MemPriority) GetClass() string {
	return b3.MEMPRIORITY
}
