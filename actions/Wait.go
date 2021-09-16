package actions

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"time"
)

/**
 * Wait a few seconds.
 *
 * @module b3
 * @class Wait
 * @extends Action
**/
type Wait struct {
	core.Action
	waitTime int64
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
func (w *Wait) Initialize(setting *config.BTNodeCfg) {
	w.Action.Initialize(setting)
	w.waitTime = setting.GetPropertyAsInt64("milliseconds")
}

/**
 * Tick method.
 * @method tick
 * @param {Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (w *Wait) OnTick(tick core.Ticker) b3.Status {
	time.Sleep(time.Duration(w.waitTime) * time.Millisecond)
	return b3.SUCCESS
}
