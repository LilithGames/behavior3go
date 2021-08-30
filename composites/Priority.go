package composites

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Priority struct {
	core.Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (p *Priority) OnTick(tick *core.Tick) b3.Status {
	for i := 0; i < p.GetChildCount(); i++ {
		var status = p.GetChild(i).Execute(tick)
		if status != b3.FAILURE {
			return status
		}
	}
	return b3.FAILURE
}

func (p *Priority) GetClass() string {
	return b3.PRIORITY
}
