package decorators

import (

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
)

/**
 * Repeater is a decorator that repeats the tick signal until the child node
 * return `RUNNING` or `ERROR`. Optionally, a maximum number of repetitions
 * can be defined.
 *
 * @module b3
 * @class Repeater
 * @extends Decorator
**/
type Repeater struct {
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
func (r *Repeater) Initialize(setting *config.BTNodeCfg) {
	r.Decorator.Initialize(setting)
	r.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if r.maxLoop < 1 {
		panic("maxLoop parameter in MaxTime decorator is an obligatory parameter")
	}
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (r *Repeater) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("i", 0, tick.GetTree().GetID(), r.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (r *Repeater) OnTick(tick *core.Tick) b3.Status {
	if r.GetChild() == nil {
		return b3.ERROR
	}
	var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), r.GetID())
	var status = b3.SUCCESS
	for r.maxLoop < 0 || i < r.maxLoop {
		status = r.GetChild().Execute(tick)
		if status == b3.SUCCESS || status == b3.FAILURE {
			i++
		} else {
			break
		}
	}
	tick.Blackboard.Set("i", i, tick.GetTree().GetID(), r.GetID())
	return status
}
