package actions

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Error struct {
	core.Action
}

func (e *Error) OnTick(tick core.Ticker) b3.Status {
	return b3.ERROR
}
