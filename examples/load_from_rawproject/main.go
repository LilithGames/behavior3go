/*
从原生工程文件加载
*/
package main

import (
	"fmt"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/examples/share"
	"github.com/magicsea/behavior3go/loader"
)

func main() {
	projectConfig, ok := config.LoadRawProjectCfg("example.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	//自定义节点注册
	maps := core.NewRegisterStructMaps()
	maps.Register("Log", new(share.LogTest))

	var firstTree *core.BehaviorTree
	//载入
	for _, v := range projectConfig.Data.Trees {
		tree := loader.CreateBevTreeFromConfig(&v, maps)
		tree.Print()
		if firstTree == nil {
			firstTree = tree
		}
	}

	//输入板
	board := core.NewBlackboard()
	//循环每一帧
	for i := 0; i < 5; i++ {
		firstTree.Tick(i, board)
	}
}
