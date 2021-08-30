package core

import (
	"fmt"
	"reflect"
	"sync"
)
/**
 * The Blackboard is the memory structure required by `BehaviorTree` and its
 * nodes. It only have 2 public methods: `set` and `get`. These methods works
 * in 3 different contexts: global, per tree, and per node per tree.
 *
 * Suppose you have two different trees controlling a single object with a
 * single blackboard, then:
 *
 * - In the global context, all nodes will access the stored information.
 * - In per tree context, only nodes sharing the same tree share the stored
 *   information.
 * - In per node per tree context, the information stored in the blackboard
 *   can only be accessed by the same node that wrote the data.
 *
 * The context is selected indirectly by the parameters provided to these
 * methods, for example:
 *
 *     // getting/setting variable in global context
 *     blackboard.set('testKey', 'value');
 *     var value = blackboard.get('testKey');
 *
 *     // getting/setting variable in per tree context
 *     blackboard.set('testKey', 'value', tree.id);
 *     var value = blackboard.get('testKey', tree.id);
 *
 *     // getting/setting variable in per node per tree context
 *     blackboard.set('testKey', 'value', tree.id, node.id);
 *     var value = blackboard.get('testKey', tree.id, node.id);
 *
 * Note: Internally, the blackboard store these memories in different
 * objects, being the global on `_baseMemory`, the per tree on `_treeMemory`
 * and the per node per tree dynamically create inside the per tree memory
 * (it is accessed via `_treeMemory[id].nodeMemory`). Avoid to use these
 * variables manually, use `get` and `set` instead.
 *
 * @module b3
 * @class Blackboard
**/
//------------------------TreeData-------------------------
type TreeData struct {
	NodeMemory     *Memory
	OpenNodes      []IBaseNode
	TraversalDepth int
	TraversalCycle int
}

func NewTreeData() *TreeData {
	return &TreeData{NewMemory(), make([]IBaseNode, 0), 0, 0}
}

//------------------------Memory-------------------------
type Memory struct {
	memory *sync.Map
}

func NewMemory() *Memory {
	return &Memory{memory: &sync.Map{}}
}

func (m *Memory) Get(key string) interface{} {
	rs, ok := m.memory.Load(key)
	if ok {
		return rs
	}
	return nil
}

func (m *Memory) Set(key string, val interface{}) {
	m.memory.Store(key, val)
}

func (m *Memory) Remove(key string) {
	m.memory.Delete(key)
}

//------------------------TreeMemory-------------------------
type TreeMemory struct {
	memory *Memory
	treeData   *TreeData
	nodeMemory *sync.Map
}

func NewTreeMemory() *TreeMemory {
	return &TreeMemory{NewMemory(), NewTreeData(), &sync.Map{}}
}

//------------------------Blackboard-------------------------
type Blackboard struct {
	baseMemory *Memory
	treeMemory *sync.Map
}

func NewBlackboard() *Blackboard {
	p := &Blackboard{}
	p.Initialize()
	return p
}

func (b *Blackboard) Initialize() {
	b.baseMemory = NewMemory()
	b.treeMemory = &sync.Map{}
}

/**
 * Internal method to retrieve the tree context memory. If the memory does
 * not exist, this method creates it.
 *
 * @method _getTreeMemory
 * @param {string} treeScope The id of the tree in scope.
 * @return {Object} The tree memory.
 * @protected
**/
func (b *Blackboard) _getTreeMemory(treeScope string) *TreeMemory {
	if rs, ok := b.treeMemory.Load(treeScope); ok {
		return rs.(*TreeMemory)
	}
	tm := NewTreeMemory()
	b.treeMemory.Store(treeScope, tm)
	return tm
}

/**
 * Internal method to retrieve the node context memory, given the tree
 * memory. If the memory does not exist, this method creates is.
 *
 * @method _getNodeMemory
 * @param {String} treeMemory the tree memory.
 * @param {String} nodeScope The id of the node in scope.
 * @return {Object} The node memory.
 * @protected
**/
func (b *Blackboard) _getNodeMemory(treeMemory *TreeMemory, nodeScope string) *Memory {
	if rs, ok := treeMemory.nodeMemory.Load(nodeScope); ok {
		return rs.(*Memory)
	}
	memory := NewMemory()
	treeMemory.nodeMemory.Store(nodeScope, memory)
	return memory
}

/**
 * Internal method to retrieve the context memory. If treeScope and
 * nodeScope are provided, this method returns the per node per tree
 * memory. If only the treeScope is provided, it returns the per tree
 * memory. If no parameter is provided, it returns the global memory.
 * Notice that, if only nodeScope is provided, this method will still
 * return the global memory.
 *
 * @method _getMemory
 * @param {String} treeScope The id of the tree scope.
 * @param {String} nodeScope The id of the node scope.
 * @return {Object} A memory object.
 * @protected
**/
func (b *Blackboard) _getMemory(treeScope, nodeScope string) *Memory {
	var memory = b.baseMemory

	if len(treeScope) > 0 {
		treeMem := b._getTreeMemory(treeScope)
		memory = treeMem.memory
		if len(nodeScope) > 0 {
			memory = b._getNodeMemory(treeMem, nodeScope)
		}
	}

	return memory
}

/**
 * Stores a value in the blackboard. If treeScope and nodeScope are
 * provided, this method will save the value into the per node per tree
 * memory. If only the treeScope is provided, it will save the value into
 * the per tree memory. If no parameter is provided, this method will save
 * the value into the global memory. Notice that, if only nodeScope is
 * provided (but treeScope not), this method will still save the value into
 * the global memory.
 *
 * @method set
 * @param {String} key The key to be stored.
 * @param {String} value The value to be stored.
 * @param {String} treeScope The tree id if accessing the tree or node
 *                           memory.
 * @param {String} nodeScope The node id if accessing the node memory.
**/
func (b *Blackboard) Set(key string, value interface{}, treeScope, nodeScope string) {
	var memory = b._getMemory(treeScope, nodeScope)
	memory.Set(key, value)
}

func (b *Blackboard) SetMem(key string, value interface{}) {
	var memory = b._getMemory("", "")
	memory.Set(key, value)
}

func (b *Blackboard) Remove(key string) {
	var memory = b._getMemory("", "")
	memory.Remove(key)
}
func (b *Blackboard) SetTree(key string, value interface{}, treeScope string) {
	var memory = b._getMemory(treeScope, "")
	memory.Set(key, value)
}
func (b *Blackboard) _getTreeData(treeScope string) *TreeData {
	treeMem := b._getTreeMemory(treeScope)
	return treeMem.treeData
}

/**
 * Retrieves a value in the blackboard. If treeScope and nodeScope are
 * provided, this method will retrieve the value from the per node per tree
 * memory. If only the treeScope is provided, it will retrieve the value
 * from the per tree memory. If no parameter is provided, this method will
 * retrieve from the global memory. If only nodeScope is provided (but
 * treeScope not), this method will still try to retrieve from the global
 * memory.
 *
 * @method get
 * @param {String} key The key to be retrieved.
 * @param {String} treeScope The tree id if accessing the tree or node
 *                           memory.
 * @param {String} nodeScope The node id if accessing the node memory.
 * @return {Object} The value stored or undefined.
**/
func (b *Blackboard) Get(key, treeScope, nodeScope string) interface{} {
	memory := b._getMemory(treeScope, nodeScope)
	return memory.Get(key)
}
func (b *Blackboard) GetMem(key string) interface{} {
	memory := b._getMemory("","")
	return memory.Get(key)
}
func (b *Blackboard) GetFloat64(key, treeScope, nodeScope string) float64 {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return v.(float64)
}
func (b *Blackboard) GetBool(key, treeScope, nodeScope string) bool {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return false
	}
	return v.(bool)
}
func (b *Blackboard) GetInt(key, treeScope, nodeScope string) int {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return v.(int)
}
func (b *Blackboard) GetInt64(key, treeScope, nodeScope string) int64 {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return v.(int64)
}
func (b *Blackboard) GetUInt64(key, treeScope, nodeScope string) uint64 {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return v.(uint64)
}

func (b *Blackboard) GetInt64Safe(key, treeScope, nodeScope string) int64 {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return ReadNumberToInt64(v)
}
func (b *Blackboard) GetUInt64Safe(key, treeScope, nodeScope string) uint64 {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return ReadNumberToUInt64(v)
}

func (b *Blackboard) GetInt32(key, treeScope, nodeScope string) int32 {
	v := b.Get(key, treeScope, nodeScope)
	if v == nil {
		return 0
	}
	return v.(int32)
}

func ReadNumberToInt64(v interface{})  int64 {
	var ret int64
	switch tvalue := v.(type) {
	case uint64:
		ret = int64(tvalue)
	default:
		panic(fmt.Sprintf("错误的类型转成Int64 %v:%+v", reflect.TypeOf(v), v))
	}

	return ret
}

func ReadNumberToUInt64(v interface{}) uint64 {
	var ret uint64
	switch tvalue := v.(type) {
	case int64:
		ret = uint64(tvalue)
	default:
		panic(fmt.Sprintf("错误的类型转成UInt64 %v:%+v", reflect.TypeOf(v), v))
	}
	return ret
}