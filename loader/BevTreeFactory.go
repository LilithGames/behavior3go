package loader

import (
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/actions"
	"github.com/magicsea/behavior3go/composites"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/decorators"
)

func createBaseStructMaps() *b3.RegisterStructMaps {
	st := b3.NewRegisterStructMaps()
	// actions
	st.Register("Error", &actions.Error{})
	st.Register("Failer", &actions.Failer{})
	st.Register("Runner", &actions.Runner{})
	st.Register("Succeeder", &actions.Succeeder{})
	st.Register("Wait", &actions.Wait{})
	st.Register("Log", &actions.Log{})
	// composites
	st.Register("MemPriority", &composites.MemPriority{})
	st.Register("MemSequence", &composites.MemSequence{})
	st.Register("Priority", &composites.Priority{})
	st.Register("Sequence", &composites.Sequence{})

	// decorators
	st.Register("Inverter", &decorators.Inverter{})
	st.Register("Limiter", &decorators.Limiter{})
	st.Register("MaxTime", &decorators.MaxTime{})
	st.Register("Repeater", &decorators.Repeater{})
	st.Register("RepeatUntilFailure", &decorators.RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &decorators.RepeatUntilSuccess{})
	return st
}

func CreateBevTreeFromConfig(config *config.BTTreeCfg, extMap *b3.RegisterStructMaps) *core.BehaviorTree {
	baseMaps := createBaseStructMaps()
	tree := core.NewBeTree()
	tree.Load(config, baseMaps, extMap)
	return tree
}
