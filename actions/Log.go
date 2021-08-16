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

func (this *Log) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *Log) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("log:", this.info)
	return b3.SUCCESS
}
