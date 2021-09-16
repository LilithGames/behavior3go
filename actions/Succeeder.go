package actions

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Succeeder struct {
	core.Action
}

func (s *Succeeder) OnTick(tick core.Ticker) b3.Status {
	return b3.SUCCESS
}
