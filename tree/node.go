package tree

type Node struct {
	Num         int
	Left, Right *Node
}

//使用new关键字创建对象
func CreateNewInstance() *Node {
	return new(Node)
}
func CreateNewInst2(arg1 int) *Node {
	return &Node{Num: arg1}
}

func (root *Node)SetNum(agr1 int)(){
	root.Num = agr1
}

