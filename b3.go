package behavior3go

// b3 define
const (
	VERSION = "0.2.0"

	// Node categories
	COMPOSITE    = "composite"
	DECORATOR    = "decorator"
	ACTION       = "action"
	CONDITION    = "condition"
	BASE         = "BaseNode"
	PARALLEL     = "Parallel"
	SUBSCRIPTION = "Subscription"
	SUBSCRIBER   = "Subscriber"
	SEQUENCE     = "Sequence"
	PRIORITY     = "Priority"
	MEMSEQUENCE  = "MemSequence"
	MEMPRIORITY  = "MemPriority"
)

// Returning status
type Status uint8

const (
	SUCCESS Status = 1
	FAILURE Status = 2
	RUNNING Status = 3
	ERROR   Status = 4
)
