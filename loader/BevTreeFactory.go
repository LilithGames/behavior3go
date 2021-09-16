package loader

import (
	"fmt"
	"github.com/magicsea/behavior3go/actions"
	"github.com/magicsea/behavior3go/composites"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/decorators"
)

func createBaseFactoryMaps() map[string]core.NodeCreator {
	result := make(map[string]core.NodeCreator)
	result["Error"] = func() core.IBaseNode {
		return &actions.Error{}
	}
	result["Failer"] = func() core.IBaseNode {
		return &actions.Failer{}
	}
	result["Runner"] = func() core.IBaseNode {
		return &actions.Runner{}
	}
	result["Succeeder"] = func() core.IBaseNode {
		return &actions.Succeeder{}
	}
	result["Wait"] = func() core.IBaseNode {
		return &actions.Wait{}
	}
	result["Log"] = func() core.IBaseNode {
		return &actions.Log{}
	}
	result["MemPriority"] = func() core.IBaseNode {
		return &composites.MemPriority{}
	}
	result["MemSequence"] = func() core.IBaseNode {
		return &composites.MemSequence{}
	}
	result["Priority"] = func() core.IBaseNode {
		return &composites.Priority{}
	}
	result["Sequence"] = func() core.IBaseNode {
		return &composites.Sequence{}
	}
	result["Parallel"] = func() core.IBaseNode {
		return &composites.Parallel{}
	}
	result["Inverter"] = func() core.IBaseNode {
		return &decorators.Inverter{}
	}
	result["Limiter"] = func() core.IBaseNode {
		return &decorators.Limiter{}
	}
	result["MaxTime"] = func() core.IBaseNode {
		return &decorators.MaxTime{}
	}
	result["Repeater"] = func() core.IBaseNode {
		return &decorators.Repeater{}
	}
	result["RepeatUntilFailure"] = func() core.IBaseNode {
		return &decorators.RepeatUntilFailure{}
	}
	result["RepeatUntilSuccess"] = func() core.IBaseNode {
		return &decorators.RepeatUntilSuccess{}
	}
	return result
}

func CreateBevTreeFromConfig(config *config.BTTreeCfg, extMap *core.RegisterStructMaps) *core.BehaviorTree {
	baseMaps := createBaseFactoryMaps()
	tree := core.NewBeTree()
	tree.Load(config, baseMaps, extMap)
	return tree
}

// Check Tree Nodes
func CheckTreeComplete(trees []config.BTTreeCfg, extMap *core.RegisterStructMaps) error {
	baseMap := createBaseFactoryMaps()
	for _, tree := range trees {
		for _, nodeCfg := range tree.Nodes {
			var exist bool
			name := nodeCfg.Name
			if extMap.CheckNode(name) {
				exist = true
			}
			if _, ok := baseMap[name]; ok {
				exist = true
			}
			if !exist && nodeCfg.Category != "tree" {
				return fmt.Errorf("not found node %s", name)
			}
		}
	}
	return nil
}
