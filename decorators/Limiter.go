package decorators

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
)

/**
 * This decorator limit the number of times its child can be called. After a
 * certain number of times, the Limiter decorator returns `FAILURE` without
 * executing the child.
 *
 * @module b3
 * @class Limiter
 * @extends Decorator
**/
type Limiter struct {
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
func (l *Limiter) Initialize(setting *config.BTNodeCfg) {
	l.Decorator.Initialize(setting)
	l.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if l.maxLoop < 1 {
		panic("maxLoop parameter in MaxTime decorator is an obligatory parameter")
	}
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (l *Limiter) OnTick(tick core.Ticker) b3.Status {
	if l.GetChild() == nil {
		return b3.ERROR
	}
	var i = tick.Blackboard().GetInt("i", tick.GetTree().GetID(), l.GetID())
	if i < l.maxLoop {
		var status = l.GetChild().Execute(tick)
		if status == b3.SUCCESS || status == b3.FAILURE {
			tick.Blackboard().Set("i", i+1, tick.GetTree().GetID(), l.GetID())
		}
		return status
	}

	return b3.FAILURE
}
