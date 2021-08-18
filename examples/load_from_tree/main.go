/*
从导出的树文件加载
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
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}
	//自定义节点注册
	maps := core.NewRegisterStructMaps()
	maps.Register("Log", new(share.LogTest))

	//载入
	tree := loader.CreateBevTreeFromConfig(treeConfig, maps)
	tree.Print()

	//输入板
	board := core.NewBlackboard()
	//循环每一帧
	for i := 0; i < 5; i++ {
		tree.Tick(i, board)
	}
}
