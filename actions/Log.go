package actions

import (
	"fmt"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
)

type Log struct {
	core.Action
	info string
}

func (l *Log) Initialize(setting *config.BTNodeCfg) {
	l.Action.Initialize(setting)
	l.info = setting.GetPropertyAsString("info")
}

func (l *Log) OnTick(tick core.Ticker) b3.Status {
	fmt.Println("log:", l.info)
	return b3.SUCCESS
}
