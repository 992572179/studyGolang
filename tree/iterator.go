package tree

import "fmt"

func (root *Node) NodeIte() {

	if root == nil {
		return
	}
	root.Left.NodeIte()
	root.Print()
	root.Right.NodeIte()

}

func (root *Node) Print() {
	fmt.Println(root.Num)
}
