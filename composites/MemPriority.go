package composites

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"time"
)

type MemPriority struct {
	core.Composite
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (p *MemPriority) OnOpen(tick core.Ticker) {
	tick.Blackboard().Set("runningChild", 0, tick.GetTree().GetID(), p.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (p *MemPriority) OnTick(tick core.Ticker) b3.Status {
	var child = tick.Blackboard().GetInt("runningChild", tick.GetTree().GetID(), p.GetID())
	for i := child; i < p.GetChildCount(); i++ {
		var status = p.GetChild(i).Execute(tick)
		for status == b3.RUNNING {
			tick.Blackboard().Set("runningChild", i, tick.GetTree().GetID(), p.GetID())
			time.Sleep(time.Second)
			status = p.GetChild(i).Execute(tick)
		}
		if status != b3.FAILURE {
			return status
		}
	}
	return b3.FAILURE
}

func (p *MemPriority) GetClass() string {
	return b3.MEMPRIORITY
}
