package core

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
)

type IAction interface {
	IBaseNode
}

/**
 * Action is the base class for all action nodes. Thus, if you want to create
 * new custom action nodes, you need to inherit from this class. For example,
 * take a look at the Runner action:
 *
 *     var Runner = b3.Class(b3.Action, {
 *       name: 'Runner',
 *
 *       tick: function(tick) {
 *         return b3.RUNNING;
 *       }
 *     });
 *
 * @module b3
 * @class Action
 * @extends BaseNode
**/
type Action struct {
	BaseNode
	BaseWorker
}

func (a *Action) Ctor() {
	a.category = b3.ACTION
}

func (a *Action) Initialize(params *config.BTNodeCfg) {

	//a.id = b3.CreateUUID()
	a.BaseNode.Initialize(params)
	//a.BaseNode.IBaseWorker = a
	a.parameters = make(map[string]interface{})
	a.properties = make(map[string]interface{})
}

func (a *Action) GetClass() string {
	return b3.ACTION
}
