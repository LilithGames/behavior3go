package core

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
)

type IDecorator interface {
	IBaseNode
	SetChild(child IBaseNode)
	GetChild() IBaseNode
}

type Decorator struct {
	BaseNode
	BaseWorker
	child IBaseNode
}

func (d *Decorator) Ctor() {
	d.category = b3.DECORATOR
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (d *Decorator) Initialize(params *config.BTNodeCfg) {
	d.BaseNode.Initialize(params)
	//d.BaseNode.IBaseWorker = d
}

//GetChild
func (d *Decorator) GetChild() IBaseNode {
	return d.child
}

func (d *Decorator) SetChild(child IBaseNode) {
	d.child = child
}

func (d *Decorator) GetClass() string {
	return b3.DECORATOR
}
