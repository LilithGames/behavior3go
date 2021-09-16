package core

import (
	"fmt"
	b3 "github.com/magicsea/behavior3go"
)

type IBaseWorker interface {

	/**
	 * Enter method, override this to use. It is called every time a node is
	 * asked to execute, before the tick itself.
	 *
	 * @method enter
	 * @param {Tick} tick A tick instance.
	**/
	OnEnter(tick Ticker)
	/**
	 * Open method, override this to use. It is called only before the tick
	 * callback and only if the not isn't closed.
	 *
	 * Note: a node will be closed if it returned `b3.RUNNING` in the tick.
	 *
	 * @method open
	 * @param {Tick} tick A tick instance.
	**/
	OnOpen(tick Ticker)
	/**
	 * Tick method, override this to use. This method must contain the real
	 * execution of node (perform a task, call children, etc.). It is called
	 * every time a node is asked to execute.
	 *
	 * @method tick
	 * @param {Tick} tick A tick instance.
	**/
	OnTick(tick Ticker) b3.Status
	/**
	 * Close method, override this to use. This method is called after the tick
	 * callback, and only if the tick return a state different from
	 * `b3.RUNNING`.
	 *
	 * @method close
	 * @param {Tick} tick A tick instance.
	**/
	OnClose(tick Ticker)
	/**
	 * Exit method, override this to use. Called every time in the end of the
	 * execution.
	 *
	 * @method exit
	 * @param {Tick} tick A tick instance.
	**/
	OnExit(tick Ticker)
}
type BaseWorker struct {
}

/**
 * Enter method, override this to use. It is called every time a node is
 * asked to execute, before the tick itself.
 *
 * @method enter
 * @param {Tick} tick A tick instance.
**/
func (w *BaseWorker) OnEnter(tick Ticker) {

}

/**
 * Open method, override this to use. It is called only before the tick
 * callback and only if the not isn't closed.
 *
 * Note: a node will be closed if it returned `b3.RUNNING` in the tick.
 *
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (w *BaseWorker) OnOpen(tick Ticker) {

}

/**
 * Tick method, override this to use. This method must contain the real
 * execution of node (perform a task, call children, etc.). It is called
 * every time a node is asked to execute.
 *
 * @method tick
 * @param {Tick} tick A tick instance.
**/
func (w *BaseWorker) OnTick(tick Ticker) b3.Status {
	fmt.Println("tick BaseWorker")
	return b3.ERROR
}

/**
 * Close method, override this to use. This method is called after the tick
 * callback, and only if the tick return a state different from
 * `b3.RUNNING`.
 *
 * @method close
 * @param {Tick} tick A tick instance.
**/
func (w *BaseWorker) OnClose(tick Ticker) {

}

/**
 * Exit method, override this to use. Called every time in the end of the
 * execution.
 *
 * @method exit
 * @param {Tick} tick A tick instance.
**/
func (w *BaseWorker) OnExit(tick Ticker) {

}
