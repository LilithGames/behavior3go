package decorators

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

/**
 * The Inverter decorator inverts the result of the child, returning `SUCCESS`
 * for `FAILURE` and `FAILURE` for `SUCCESS`.
 *
 * @module b3
 * @class Inverter
 * @extends Decorator
**/
type Inverter struct {
	core.Decorator
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (i *Inverter) OnTick(tick core.Ticker) b3.Status {
	if i.GetChild() == nil {
		return b3.ERROR
	}

	var status = i.GetChild().Execute(tick)
	if status == b3.SUCCESS {
		status = b3.FAILURE
	} else if status == b3.FAILURE {
		status = b3.SUCCESS
	}

	return status
}
