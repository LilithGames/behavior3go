package decorators

import (
	"time"

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
type MaxTime struct {
	core.Decorator
	maxTime int64
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
func (t *MaxTime) Initialize(setting *config.BTNodeCfg) {
	t.Decorator.Initialize(setting)
	t.maxTime = setting.GetPropertyAsInt64("maxTime")
	if t.maxTime < 1 {
		panic("maxTime parameter in Limiter decorator is an obligatory parameter")
	}
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (t *MaxTime) OnOpen(tick core.Ticker) {
	var startTime = time.Now().UnixNano() / 1000000
	tick.Blackboard().Set("startTime", startTime, tick.GetTree().GetID(), t.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (t *MaxTime) OnTick(tick core.Ticker) b3.Status {
	if t.GetChild() == nil {
		return b3.ERROR
	}
	var currTime = time.Now().UnixNano() / 1000000
	var startTime int64 = tick.Blackboard().GetInt64("startTime", tick.GetTree().GetID(), t.GetID())
	var status = t.GetChild().Execute(tick)
	if currTime-startTime > t.maxTime {
		return b3.FAILURE
	}
	return status
}
