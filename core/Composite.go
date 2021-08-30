package core

import (
	"fmt"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
)

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	BaseWorker
	children []IBaseNode
}

func (c *Composite) Ctor() {
	c.category = b3.COMPOSITE
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (c *Composite) Initialize(params *config.BTNodeCfg) {
	c.BaseNode.Initialize(params)
	c.children = make([]IBaseNode, 0)
}

/**
 *
 * @method GetChildCount
 * @getChildCount
**/
func (c *Composite) GetChildCount() int {
	return len(c.children)
}

// GetChild
func (c *Composite) GetChild(index int) IBaseNode {
	return c.children[index]
}

// AddChild
func (c *Composite) AddChild(child IBaseNode) {
	c.children = append(c.children, child)
}

func (c *Composite) tick(tick *Tick) b3.Status {
	fmt.Println("tick Composite1")
	return b3.ERROR
}

func (c *Composite) GetClass() string {
	return b3.COMPOSITE
}
