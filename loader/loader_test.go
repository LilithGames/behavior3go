package loader

import (
	"fmt"
	"reflect"
	"testing"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
)

type Test struct {
	value string
}

func (test *Test) Print() {
	fmt.Println(test.value)
}

func TestExample(t *testing.T) {
	maps := createBaseStructMaps()
	if data, err := maps.New("Runner"); err != nil {
		t.Error("Error:", err, data)
	} else {
		t.Log(reflect.TypeOf(data))
	}

}

///////////////////////加载事例///////////////////////////
//自定义action节点
type LogTest struct {
	core.Action
	info string
}

func (this *LogTest) Initialize(setting *config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("logtest:", this.info)
	return b3.SUCCESS
}

func TestLoadTree(t *testing.T) {
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if ok {
		//自定义节点注册
		maps := b3.NewRegisterStructMaps()
		maps.Register("Log", new(LogTest))

		//载入
		tree := CreateBevTreeFromConfig(treeConfig, maps)
		tree.Print()

		//输入板
		board :=core.NewBlackboard()
		//循环每一帧
		for i := 0; i < 5; i++ {
			tree.Tick(i, board)
		}
	} else {
		t.Error("LoadTreeCfg err")
	}
}
