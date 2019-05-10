package intervaltree

import "fmt"

const (
	//红树
	RED bool = true
	//黑树
	BLACK bool = false
)

type INode struct {
	//区间树节点
	Cid                 int8   //课程编号
	Cname               string //课程名
	Low, High, Max      int8   //区间变量
	Icolor              bool   //区间树颜色
	Parent, Left, Right *INode //区间树指针
}

//创建新节点
func NewINode(id, Low, High, Max int8, name string, color bool) *INode {
	return &INode{Cid: id, Cname: name, Low: Low, High: High, Max: Max, Icolor: color, Left: nil, Right: nil, Parent: nil}
}

//输出节点
func (node *INode) PrintNode(itree ITree) {
	if node == itree.Nil{
		fmt.Println("Nil Node ")
		return
	}
	tstr := ""
	if node.Icolor {
		tstr = "RED"
	} else {
		tstr = "BLACK"
	}
	fmt.Printf("\n%d,\t %s,\t [%d,\t%d]\t%d,\t%s,\t",node.Cid, node.Cname, node.Low, node.High,node.Max,tstr)
	if node.Left != itree.Nil {
		fmt.Printf("Leftchild Cid = %d\t", node.Left.Cid)
	}
	if node.Right != itree.Nil {
		fmt.Printf("Rightchile Cid = %d\t", node.Right.Cid)
	}
	fmt.Printf("\n")
}
