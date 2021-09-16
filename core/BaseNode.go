package core

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
)

type IBaseWrapper interface {
	_execute(tick Ticker) b3.Status
	_enter(tick Ticker)
	_open(tick Ticker)
	_tick(tick Ticker) b3.Status
	_close(tick Ticker)
	_exit(tick Ticker)
}
type IBaseNode interface {
	IBaseWrapper
	GetID() string
	GetParentID() string
	SetTreeID(id string)
	GetTreeID() string
	Ctor()
	Initialize(params *config.BTNodeCfg)
	GetCategory() string
	Execute(tick Ticker) b3.Status
	GetName() string
	GetTitle() string
	GetParent() IBaseNode
	SetParent(node IBaseNode)
	SetBaseNodeWorker(worker IBaseWorker)
	GetBaseNodeWorker() IBaseWorker
	GetValueFromAncestor(key string, blackboard *Blackboard) interface{}
	GetClass() string
}

/**
 * The BaseNode class is used as super class to all nodes in BehaviorJS. It
 * comprises all common variables and methods that a node must have to
 * execute.
 *
 * **IMPORTANT:** Do not inherit from this class, use `b3.Composite`,
 * `b3.Decorator`, `b3.Action` or `b3.Condition`, instead.
 *
 * The attributes are specially designed to serialization of the node in a
 * JSON format. In special, the `parameters` attribute can be set into the
 * visual editor (thus, in the JSON file), and it will be used as parameter
 * on the node initialization at `BehaviorTree.load`.
 *
 * BaseNode also provide 5 callback methods, which the node implementations
 * can override. They are `enter`, `open`, `tick`, `close` and `exit`. See
 * their documentation to know more. These callbacks are called inside the
 * `_execute` method, which is called in the tree traversal.
 *
 * @module b3
 * @class BaseNode
**/
type BaseNode struct {
	IBaseWorker
	/**
	 * Node ID.
	 * @property {string} id
	 * @readonly
	**/
	id string

	/**
	 * Node name. Must be a unique identifier, preferable the same name of the
	 * class. You have to set the node name in the prototype.
	 *
	 * @property {String} name
	 * @readonly
	**/
	name string

	/**
	 * Node category. Must be `b3.COMPOSITE`, `b3.DECORATOR`, `b3.ACTION` or
	 * `b3.CONDITION`. This is defined automatically be inheriting the
	 * correspondent class.
	 *
	 * @property {CONSTANT} category
	 * @readonly
	**/
	category string

	/**
	 * Node title.
	 * @property {String} title
	 * @optional
	 * @readonly
	**/
	title string

	/**
	 * Node description.
	 * @property {String} description
	 * @optional
	 * @readonly
	**/
	description string

	/**
	 * A dictionary (key, value) describing the node parameters. Useful for
	 * defining parameter values in the visual editor. Note: this is only
	 * useful for nodes when loading trees from JSON files.
	 *
	 * **Deprecated since 0.2.0. This is too similar to the properties
	 * attribute, thus, this attribute is deprecated in favor to
	 * `properties`.**
	 *
	 * @property {Object} parameters
	 * @deprecated since 0.2.0.
	 * @readonly
	**/
	parameters map[string]interface{}

	/**
	 * A dictionary (key, value) describing the node properties. Useful for
	 * defining custom variables inside the visual editor.
	 *
	 * @property properties
	 * @type {Object}
	 * @readonly
	**/
	properties map[string]interface{}

	parent IBaseNode

	treeID string
}

func (n *BaseNode) Ctor() {
}

func (n *BaseNode) SetName(name string) {
	n.name = name
}
func (n *BaseNode) SetTitle(name string) {
	n.name = name
}

func (n *BaseNode) SetBaseNodeWorker(worker IBaseWorker) {
	n.IBaseWorker = worker
}

func (n *BaseNode) GetBaseNodeWorker() IBaseWorker {
	return n.IBaseWorker
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (n *BaseNode) Initialize(params *config.BTNodeCfg) {
	n.description = ""
	n.parameters = make(map[string]interface{})
	n.properties = make(map[string]interface{})

	n.id = params.Id // node.id;
	n.name = params.Name
	n.title = params.Title             // node.title;
	n.description = params.Description // node.description;
	n.properties = params.Properties   // node.properties;
}

func (n *BaseNode) GetCategory() string {
	return n.category
}

func (n *BaseNode) GetID() string {
	return n.id
}

func (n *BaseNode) GetParentID() string {
	if n.parent == nil {
		return ""
	}
	return n.parent.GetID()
}

func (n *BaseNode) GetName() string {
	return n.name
}

func (n *BaseNode) GetTitle() string {
	return n.title
}

func (n *BaseNode) GetParent() IBaseNode {
	return n.parent
}

func (n *BaseNode) SetParent(parent IBaseNode) {
	n.parent = parent
}

func (n *BaseNode) GetValueFromAncestor(key string, blackboard *Blackboard) interface{} {
	parent := n.GetParent()
	for {
		if parent == nil {
			return nil
		}
		if v := blackboard.Get(key, n.GetTreeID(), parent.GetID()); v != nil {
			return v
		}
		parent = parent.GetParent()
	}
}

func (n *BaseNode) GetTreeID() string {
	return n.treeID
}

func (n *BaseNode) SetTreeID(treeID string) {
	n.treeID = treeID
}

func (n *BaseNode) GetClass() string {
	return b3.BASE
}

/**
 * This is the main method to propagate the tick signal to this node. This
 * method calls all callbacks: `enter`, `open`, `tick`, `close`, and
 * `exit`. It only opens a node if it is not already open. In the same
 * way, this method only close a node if the node  returned a status
 * different of `b3.RUNNING`.
 *
 * @method _execute
 * @param {Tick} tick A tick instance.
 * @return {Constant} The tick state.
 * @protected
**/
func (n *BaseNode) _execute(tick Ticker) b3.Status {
	//fmt.Println("_execute :", n.title)
	// ENTER
	n._enter(tick)

	// OPEN
	if !tick.Blackboard().GetBool("isOpen", tick.GetTree().id, n.id) {
		n._open(tick)
	}

	// TICK
	var status = n._tick(tick)

	// CLOSE
	if status != b3.RUNNING {
		n._close(tick)
	}

	// EXIT
	n._exit(tick)

	return status
}
func (n *BaseNode) Execute(tick Ticker) b3.Status {
	return n._execute(tick)
}

/**
 * Wrapper for enter method.
 * @method _enter
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (n *BaseNode) _enter(tick Ticker) {
	tick._enterNode(n)
	n.OnEnter(tick)
}

/**
 * Wrapper for open method.
 * @method _open
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (n *BaseNode) _open(tick Ticker) {
	//fmt.Println("_open :", n.title)
	tick._openNode(n)
	tick.Blackboard().Set("isOpen", true, tick.GetTree().id, n.id)
	n.OnOpen(tick)
}

/**
 * Wrapper for tick method.
 * @method _tick
 * @param {Tick} tick A tick instance.
 * @return {Constant} A state constant.
 * @protected
**/
func (n *BaseNode) _tick(tick Ticker) b3.Status {
	//fmt.Println("_tick :", n.title)
	tick._tickNode(n)
	return n.OnTick(tick)
}

/**
 * Wrapper for close method.
 * @method _close
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (n *BaseNode) _close(tick Ticker) {
	tick._closeNode(n)
	tick.Blackboard().Set("isOpen", false, tick.GetTree().id, n.id)
	n.OnClose(tick)
}

/**
 * Wrapper for exit method.
 * @method _exit
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (n *BaseNode) _exit(tick Ticker) {
	tick._exitNode(n)
	n.OnExit(tick)
}
