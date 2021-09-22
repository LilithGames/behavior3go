package core

import (
	"fmt"
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
)

/**
 * The BehaviorTree class, as the name implies, represents the Behavior Tree
 * structure.
 *
 * There are two ways to construct a Behavior Tree: by manually setting the
 * root node, or by loading it from a data structure (which can be loaded
 * from a JSON). Both methods are shown in the examples below and better
 * explained in the user guide.
 *
 * The tick method must be called periodically, in order to send the tick
 * signal to all nodes in the tree, starting from the root. The method
 * `BehaviorTree.tick` receives a target object and a blackboard as
 * parameters. The target object can be anything: a game agent, a system, a
 * DOM object, etc. This target is not used by any piece of Behavior3JS,
 * i.e., the target object will only be used by custom nodes.
 *
 * The blackboard is obligatory and must be an instance of `Blackboard`. This
 * requirement is necessary due to the fact that neither `BehaviorTree` or
 * any node will store the execution variables in its own object (e.g., the
 * BT does not store the target, information about opened nodes or number of
 * times the tree was called). But because of this, you only need a single
 * tree instance to control multiple (maybe hundreds) objects.
 *
 * Manual construction of a Behavior Tree
 * --------------------------------------
 *
 *     var tree = new b3.BehaviorTree();
 *
 *     tree.root = new b3.Sequence({children:[
 *       new b3.Priority({children:[
 *         new MyCustomNode(),
 *         new MyCustomNode()
 *       ]}),
 *       ...
 *     ]});
 *
 *
 * Loading a Behavior Tree from data structure
 * -------------------------------------------
 *
 *     var tree = new b3.BehaviorTree();
 *
 *     tree.load({
 *       'title'       : 'Behavior Tree title'
 *       'description' : 'My description'
 *       'root'        : 'node-id-1'
 *       'nodes'       : {
 *         'node-id-1' : {
 *           'name'        : 'Priority', // this is the node type
 *           'title'       : 'Root Node',
 *           'description' : 'Description',
 *           'children'    : ['node-id-2', 'node-id-3'],
 *         },
 *         ...
 *       }
 *     })
 *
 *
 * @module b3
 * @class BehaviorTree
**/
type BehaviorTree struct {

	/**
	 * The tree id, must be unique. By default, created with `b3.createUUID`.
	 * @property {String} id
	 * @readOnly
	**/
	id string

	/**
	 * The tree title.
	 * @property {String} title
	 * @readonly
	**/
	title string

	/**
	 * Description of the tree.
	 * @property {String} description
	 * @readonly
	**/
	description string

	/**
	 * A dictionary with (key-value) properties. Useful to define custom
	 * variables in the visual editor.
	 *
	 * @property {Object} properties
	 * @readonly
	**/
	properties map[string]interface{}

	/**
	 * The reference to the root node. Must be an instance of `b3.BaseNode`.
	 * @property {BaseNode} root
	**/
	root IBaseNode

	/**
	 * The reference to the debug instance.
	 * @property {Object} debug
	**/
	debug interface{}

	dumpInfo *config.BTTreeCfg
}

func NewBeTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (t *BehaviorTree) Initialize() {
	t.id = CreateUUID()
	t.title = "The behavior tree"
	t.description = "Default description"
	t.properties = make(map[string]interface{})
	t.root = nil
	t.debug = nil
}

func (t *BehaviorTree) GetID() string {
	return t.id
}

func (t *BehaviorTree) GetTitile() string {
	return t.title
}

func (t *BehaviorTree) SetDebug(debug interface{}) {
	t.debug = debug
}

func (t *BehaviorTree) GetRoot() IBaseNode {
	return t.root
}

/**
 * This method loads a Behavior Tree from a data structure, populating this
 * object with the provided data. Notice that, the data structure must
 * follow the format specified by Behavior3JS. Consult the guide to know
 * more about this format.
 *
 * You probably want to use custom nodes in your BTs, thus, you need to
 * provide the `names` object, in which this method can find the nodes by
 * `names[NODE_NAME]`. This variable can be a namespace or a dictionary,
 * as long as this method can find the node by its name, for example:
 *
 *     //json
 *     ...
 *     'node1': {
 *       'name': MyCustomNode,
 *       'title': ...
 *     }
 *     ...
 *
 *     //code
 *     var bt = new b3.BehaviorTree();
 *     bt.load(data, {'MyCustomNode':MyCustomNode})
 *
 *
 * @method load
 * @param {Object} data The data structure representing a Behavior Tree.
 * @param {Object} [names] A namespace or dict containing custom nodes.
**/
func (t *BehaviorTree) Load(data *config.BTTreeCfg, maps map[string]NodeCreator, extMaps *RegisterStructMaps) {
	t.title = data.Title             // || t.title;
	t.description = data.Description // || t.description;
	t.properties = data.Properties   // || t.properties;
	t.dumpInfo = data
	nodes := make(map[string]IBaseNode)
	// Create the node list (without connection between them)
	for id, s := range data.Nodes {
		spec := &s
		var node IBaseNode

		if spec.Category == "tree" {
			node = new(SubTree)
		} else {
			if extMaps != nil && extMaps.CheckNode(spec.Name) {
				node = extMaps.GetNode(spec.Name)()
			} else if creator, ok := maps[spec.Name]; ok {
				node = creator()
			}
		}

		if node == nil {
			// Invalid node name
			panic("BehaviorTree.load: Invalid node name:" + spec.Name + ",title:" + spec.Title)
		}

		node.Ctor()
		node.Initialize(spec)
		node.SetBaseNodeWorker(node.(IBaseWorker))
		node.SetTreeID(data.ID)
		nodes[id] = node
	}

	// Connect the nodes
	for id, spec := range data.Nodes {
		node := nodes[id]
		if node.GetCategory() == b3.COMPOSITE && spec.Children != nil {
			for i := 0; i < len(spec.Children); i++ {
				cid := spec.Children[i]
				child := nodes[cid]
				comp := node.(IComposite)
				comp.AddChild(child)
				child.SetParent(node)
			}
		} else if node.GetCategory() == b3.DECORATOR && len(spec.Child) > 0 {
			dec := node.(IDecorator)
			child := nodes[spec.Child]
			dec.SetChild(child)
			child.SetParent(dec)
		}
	}
	t.root = nodes[data.Root]
}

/**
 * This method dump the current BT into a data structure.
 *
 * Note: This method does not record the current node parameters. Thus,
 * it may not be compatible with load for now.
 *
 * @method dump
 * @return {Object} A data object representing this tree.
**/
func (t *BehaviorTree) dump() *config.BTTreeCfg {
	return t.dumpInfo
}

/**
 * Propagates the tick signal through the tree, starting from the root.
 *
 * This method receives a target object of any type (Object, Array,
 * DOMElement, whatever) and a `Blackboard` instance. The target object has
 * no use at all for all Behavior3JS components, but surely is important
 * for custom nodes. The blackboard instance is used by the tree and nodes
 * to store execution variables (e.g., last node running) and is obligatory
 * to be a `Blackboard` instance (or an object with the same interface).
 *
 * Internally, this method creates a Tick object, which will store the
 * target and the blackboard objects.
 *
 * Note: BehaviorTree stores a list of open nodes from last tick, if these
 * nodes weren't called after the current tick, this method will close them
 * automatically.
 *
 * @method tick
 * @param {Object} target A target object.
 * @param {Blackboard} blackboard An instance of blackboard object.
 * @return {Constant} The tick signal state.
**/
func (t *BehaviorTree) Tick(tick Ticker, blackboard *Blackboard) b3.Status {
	if blackboard == nil {
		panic("The blackboard parameter is obligatory and must be an instance of b3.Blackboard")
	}

	/* CREATE A TICK OBJECT */
	tick.setTree(t)
	tick.setDebug(t.debug)
	tick.setBlackboard(blackboard)

	/* TICK NODE */
	var state = t.root._execute(tick)

	/* CLOSE NODES FROM LAST TICK, IF NEEDED */
	var lastOpenNodes = blackboard._getTreeData(t.id).OpenNodes
	var currOpenNodes []IBaseNode
	currOpenNodes = append(currOpenNodes, tick.openNodes()...)

	// does not close if it is still open in t tick
	var start = 0
	for i := 0; i < MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i + 1
		if lastOpenNodes[i] != currOpenNodes[i] {
			break
		}
	}

	// close the nodes
	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i]._close(tick)
	}

	/* POPULATE BLACKBOARD */
	blackboard._getTreeData(t.id).OpenNodes = currOpenNodes
	blackboard.SetTree("nodeCount", tick.nodeCount(), t.id)

	return state
}

func (t *BehaviorTree) Print() {
	printNode(t.root, 0)
}

func printNode(root IBaseNode, blk int) {

	for i := 0; i < blk; i++ {
		fmt.Print(" ") //缩进
	}
	fmt.Print("|—", root.GetTitle())

	if root.GetCategory() == b3.DECORATOR {
		dec := root.(IDecorator)
		if dec.GetChild() != nil {
			printNode(dec.GetChild(), blk+3)
		}
	}

	fmt.Println("")
	if root.GetCategory() == b3.COMPOSITE {
		comp := root.(IComposite)
		if comp.GetChildCount() > 0 {
			for i := 0; i < comp.GetChildCount(); i++ {
				printNode(comp.GetChild(i), blk+3)
			}
		}
	}
}
