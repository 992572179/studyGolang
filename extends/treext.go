package extends

import "github.com/study-golang/tree"

type NodeHelper struct {
	node *tree.Node
}

func (inst *NodeHelper)nodeIte()  {
	if inst == nil{
		return
	}
	inst.node.Left.NodeIte()
	inst.node.Right.NodeIte()
	inst.node.Print()
}


