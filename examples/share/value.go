package share

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
)

//自定义action节点
type SetValue struct {
	core.Action
	value int
	key   string
}

func (this *SetValue) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *SetValue) OnTick(tick core.Ticker) b3.Status {
	tick.GetBlackBoard().SetMem(this.key, this.value)
	return b3.SUCCESS
}

//自定义action节点
type IsValue struct {
	core.Condition
	value int
	key   string
}

func (this *IsValue) Initialize(setting *config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *IsValue) OnTick(tick core.Ticker) b3.Status {
	v := tick.GetBlackBoard().GetInt(this.key, "", "")
	if v == this.value {
		return b3.SUCCESS
	}
	return b3.FAILURE
}
