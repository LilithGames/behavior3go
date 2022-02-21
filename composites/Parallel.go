package composites

import (
	"context"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

type Parallel struct {
	core.Composite
	cancel context.CancelFunc
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (p *Parallel) OnOpen(tick core.Ticker) {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel
	tick.Blackboard().Set("cancelCtx", ctx, p.GetTreeID(), p.GetID())
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
	ctx := tick.Blackboard().Get("cancelCtx", p.GetTreeID(), p.GetID()).(context.Context)
	for i := 0; i < childNum; i++ {
		child := p.GetChild(i)
		nt := tick.TearTick()
		go func() {
			status := child.Execute(nt)
			for status == b3.RUNNING {
				select {
				case <-ctx.Done():
					status = b3.SUCCESS
				default:
					status = child.Execute(nt)
				}
			}
			rs <- status
		}()
	}
	var finish int
	<- rs
	finish++
	p.cancel()
	for finish < childNum {
		<- rs
		finish++
	}
	return b3.SUCCESS
}

func (p *Parallel) GetClass() string {
	return b3.PARALLEL
}
