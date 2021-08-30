package actions

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Failer struct {
	core.Action
}

func (f *Failer) OnTick(tick *core.Tick) b3.Status {
	return b3.FAILURE
}
