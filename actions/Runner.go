package actions

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Runner struct {
	core.Action
}

func (r *Runner) OnTick(tick *core.Tick) b3.Status {
	return b3.RUNNING
}
