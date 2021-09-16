package composites

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Sequence struct {
	core.Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (s *Sequence) OnTick(tick core.Ticker) b3.Status {
	for i := 0; i < s.GetChildCount(); i++ {
		var status = s.GetChild(i).Execute(tick)
		if status != b3.SUCCESS {
			return status
		}
	}
	return b3.SUCCESS
}

func (s *Sequence) GetClass() string {
	return b3.SEQUENCE
}
