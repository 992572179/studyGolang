package main

import (
	"github.com/study-golang/tree"
	"fmt"
)

func main() {

	var root = new(tree.Node)

	inst3 := tree.CreateNewInstance()
	inst3.SetNum(5)
	fmt.Println(inst3)

	inst4 := tree.CreateNewInst2(1)
	fmt.Println(inst4)

	root.Left = inst3
	root.Right = inst4
	fmt.Println("-->root iterator...")
	root.NodeIte()

}
