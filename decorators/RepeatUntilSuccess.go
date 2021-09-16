package decorators

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
)

/**
 * The MaxTime decorator limits the maximum time the node child can execute.
 * Notice that it does not interrupt the execution itself (i.e., the child
 * must be non-preemptive), it only interrupts the node after a `RUNNING`
 * status.
 *
 * @module b3
 * @class MaxTime
 * @extends Decorator
**/
type RepeatUntilSuccess struct {
	core.Decorator
	maxLoop int
}

/**
 * Initialization method.
 *
 * Settings parameters:
 *
 * - **milliseconds** (*Integer*) Maximum time, in milliseconds, a child
 *                                can execute.
 *
 * @method Initialize
 * @param {Object} settings Object with parameters.
 * @construCtor
**/
func (s *RepeatUntilSuccess) Initialize(setting *config.BTNodeCfg) {
	s.Decorator.Initialize(setting)
	s.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if s.maxLoop < 1 {
		panic("maxLoop parameter in MaxTime decorator is an obligatory parameter")
	}
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (s *RepeatUntilSuccess) OnOpen(tick core.Ticker) {
	tick.Blackboard().Set("i", 0, tick.GetTree().GetID(), s.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (s *RepeatUntilSuccess) OnTick(tick core.Ticker) b3.Status {
	if s.GetChild() == nil {
		return b3.ERROR
	}
	var i = tick.Blackboard().GetInt("i", tick.GetTree().GetID(), s.GetID())
	var status = b3.ERROR
	for s.maxLoop < 0 || i < s.maxLoop {
		status = s.GetChild().Execute(tick)
		if status == b3.FAILURE {
			i++
		} else {
			break
		}
	}

	tick.Blackboard().Set("i", i, tick.GetTree().GetID(), s.GetID())
	return status
}
