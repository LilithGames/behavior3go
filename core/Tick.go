package core

type Ticker interface {
	Initialize()
	GetTree() *BehaviorTree
	GetLastSubTree() *SubTree
	Blackboard() *Blackboard
	Tear(ticker Ticker)
	TearTick() Ticker
	_enterNode(node IBaseNode)
	_openNode(node *BaseNode)
	_tickNode(node *BaseNode)
	_closeNode(node *BaseNode)
	_exitNode(node *BaseNode)
	nodeCount() int
	openNodes() []IBaseNode
	pushSubtreeNode(node *SubTree)
	popSubtreeNode()
	setTree(tree *BehaviorTree)
	setBlackboard(blackboard *Blackboard)
	setDebug(debug interface{})
}

/**
 * A new Tick object is instantiated every tick by BehaviorTree. It is passed
 * as parameter to the nodes through the tree during the traversal.
 *
 * The role of the Tick class is to store the instances of tree, debug,
 * target and blackboard. So, all nodes can access these informations.
 *
 * For internal uses, the Tick also is useful to store the open node after
 * the tick signal, in order to let `BehaviorTree` to keep track and close
 * them when necessary.
 *
 * This class also makes a bridge between nodes and the debug, passing the
 * node state to the debug if the last is provided.
 *
 * @module b3
 * @class Tick
**/
type Tick struct {
	Ticker
	/**
	 * The tree reference.
	 * @property {b3.BehaviorTree} tree
	 * @readOnly
	**/
	tree *BehaviorTree
	/**
	 * The debug reference.
	 * @property {Object} debug
	 * @readOnly
	 */
	debug interface{}
	/**
	 * The blackboard reference.
	 * @property {b3.Blackboard} blackboard
	 * @readOnly
	**/
	blackboard *Blackboard
	/**
	 * The list of open nodes. Update during the tree traversal.
	 * @property {Array} _openNodes
	 * @protected
	 * @readOnly
	**/
	_openNodes []IBaseNode

	/**
	 * The list of open subtree node.
	 * push subtree node before execute subtree.
	 * pop subtree node after execute subtree.
	**/
	_openSubtreeNodes []*SubTree

	/**
	 * The number of nodes entered during the tick. Update during the tree
	 * traversal.
	 *
	 * @property {Integer} _nodeCount
	 * @protected
	 * @readOnly
	**/
	_nodeCount int
}

func NewTick() *Tick {
	tick := &Tick{}
	tick.Initialize()
	return tick
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (t *Tick) Initialize() {
	// set by BehaviorTree
	t.tree = nil
	t.debug = nil
	t.blackboard = nil

	// updated during the tick signal
	t._openNodes = nil
	t._openSubtreeNodes = nil
	t._nodeCount = 0
}

func (t *Tick) GetTree() *BehaviorTree {
	return t.tree
}

/**
 * Called when entering a node (called by BaseNode).
 * @method _enterNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (t *Tick) _enterNode(node IBaseNode) {
	t._nodeCount++
	t._openNodes = append(t._openNodes, node)

	// TODO: call debug here
}

/**
 * Callback when opening a node (called by BaseNode).
 * @method _openNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (t *Tick) _openNode(node *BaseNode) {
	// TODO: call debug here
}

/**
 * Callback when ticking a node (called by BaseNode).
 * @method _tickNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (t *Tick) _tickNode(node *BaseNode) {
	// TODO: call debug here
}

/**
 * Callback when closing a node (called by BaseNode).
 * @method _closeNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (t *Tick) _closeNode(node *BaseNode) {
	// TODO: call debug here
	ulen := len(t._openNodes)
	if ulen > 0 {
		t._openNodes = t._openNodes[:ulen-1]
	}

}

func (t *Tick) pushSubtreeNode(node *SubTree) {
	t._openSubtreeNodes = append(t._openSubtreeNodes, node)
}
func (t *Tick) popSubtreeNode() {
	ulen := len(t._openSubtreeNodes)
	if ulen > 0 {
		t._openSubtreeNodes = t._openSubtreeNodes[:ulen-1]
	}
}

/**
 * return top subtree node.
 * return nil when it is runing at major tree
 *
**/
func (t *Tick) GetLastSubTree() *SubTree {
	ulen := len(t._openSubtreeNodes)
	if ulen > 0 {
		return t._openSubtreeNodes[ulen-1]
	}
	return nil
}

/**
 * Callback when exiting a node (called by BaseNode).
 * @method _exitNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (t *Tick) _exitNode(node *BaseNode) {
	// TODO: call debug here
}

func (t *Tick) Blackboard() *Blackboard {
	return t.blackboard
}

func (t *Tick) setTree(tree *BehaviorTree) {
	t.tree = tree
}

func (t *Tick) nodeCount() int {
	return t._nodeCount
}

func (t *Tick) openNodes() []IBaseNode {
	return t._openNodes
}

func (t *Tick) setBlackboard(blackboard *Blackboard) {
	t.blackboard = blackboard
}

func (t *Tick) setDebug(debug interface{}) {
	t.debug = debug
}

func (t *Tick) Tear(ticker Ticker) {
	tick := ticker.(*Tick)
	tick.blackboard = t.blackboard
	tick._openNodes = append(tick._openNodes, t._openNodes...)
	tick._nodeCount = t._nodeCount
	tick.debug = t.debug
	tick.tree = t.tree
	tick._openSubtreeNodes = append(tick._openSubtreeNodes, t._openSubtreeNodes...)
}

func (t *Tick) TearTick() Ticker {
	tick := NewTick()
	t.Tear(tick)
	return tick
}
