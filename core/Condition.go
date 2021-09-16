package core

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
)

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (c *Condition) Ctor() {
	c.category = b3.CONDITION
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (c *Condition) Initialize(params *config.BTNodeCfg) {
	c.BaseNode.Initialize(params)
}

func (c *Condition) GetClass() string {
	return b3.CONDITION
}
