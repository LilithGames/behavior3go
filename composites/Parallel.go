package composites

import (
	"context"
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"time"
)

type Parallel struct {
	core.Composite
	cancel context.CancelFunc
	subNum int32
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (p *Parallel) OnOpen(tick core.Ticker) {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel
	p.subNum = 0
	tick.Blackboard().Set("cancelCtx", ctx, p.GetTreeID(), p.GetID())
	tick.Blackboard().Set("subSum", &p.subNum, p.GetTreeID(), p.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (p *Parallel) OnTick(tick core.Ticker) b3.Status {
	childNum := p.GetChildCount()
	rs := make(chan b3.Status, childNum)
	for i := 0; i < childNum; i++ {
		child := p.GetChild(i)
		go func() {
			for {
				status := child.Execute(tick.TearTick())
				if status != b3.RUNNING {
					rs <- status
					break
				}
			}
		}()
	}
	var finishCount int32
	for {
		select {
		case <-rs:
			finishCount++
			if finishCount == int32(childNum) {
				return b3.SUCCESS
			}
		case <-time.After(time.Second):
			if p.subNum+finishCount == int32(childNum) {
				p.cancel()
			}
		}
	}
}

func (p *Parallel) GetClass() string {
	return b3.PARALLEL
}
