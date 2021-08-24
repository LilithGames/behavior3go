package composites

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"sync"
)

type Parallel struct {
	core.Composite
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (p *Parallel) OnOpen(tick *core.Tick) {
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (p *Parallel) OnTick(tick *core.Tick) b3.Status {
	count := p.GetChildCount()
	group := sync.WaitGroup{}
	group.Add(count)
	rs := make(chan b3.Status, count)
	for i := 0; i < count; i++ {
		child := p.GetChild(i)
		go func() {
			rs <- child.Execute(tick)
			group.Done()
		}()
	}
	group.Wait()
	close(rs)
	for status := range rs {
		if status != b3.SUCCESS {
			return status
		}
	}
	return b3.SUCCESS
}
