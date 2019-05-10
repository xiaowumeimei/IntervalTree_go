//实现区间树
package intervaltree

import "fmt"
//区间树
type ITree struct {
	Root *INode
	Nil  *INode
}

//创建区间树
func NewITree() *ITree {
	tmp := new(ITree)
	tmp.Nil = new(INode)
	tmp.Nil.Left, tmp.Nil.Right, tmp.Nil.Parent = nil, nil, nil
	tmp.Nil.Icolor = BLACK
	tmp.Root = tmp.Nil
	return tmp
}

func (itree ITree) Minmum(node *INode) *INode {
	tmpnode := node
	for tmpnode.Left != itree.Nil {
		tmpnode = tmpnode.Left
	}
	return tmpnode
}

//返回x,y,z中的最大值
func MaxofInt8(x, y, z int8) int8{
	if x >= y {
		if x >= z {
			return x
		} else {
			return z
		}
	} else {
		if y >= z {
			return y
		} else {
			return z
		}
	}

}

//递归遍历树维护max值，调试时使用，
func (itree *ITree) MaxFixup(node *INode) int8{
	if node == itree.Nil{
		return 0
	}
	if node.Left ==itree.Nil && node.Right == itree.Nil{
		node.Max = node.High
		return node.Max
	} else {
		node.Max = MaxofInt8(node.Max,itree.MaxFixup(node.Left), itree.MaxFixup(node.Right))
		return node.Max
	}
}

//左旋
func (itree *ITree) LeftRotate(tx *INode) {
	ty := tx.Right
	tx.Right = ty.Left
	if ty.Left != itree.Nil {
		ty.Left.Parent = tx
	}
	ty.Parent = tx.Parent
	if tx.Parent == itree.Nil {
		itree.Root = ty
	} else if tx == tx.Parent.Left {
		tx.Parent.Left = ty
	} else {
		tx.Parent.Right = ty
	}
	ty.Left = tx
	tx.Parent = ty
	//左旋右旋看似花里胡哨，实际上只有tx节点和tx的子节点：ty 的max值需要维护，其余节点及其子树的max值都满足区间树定义
	tx.Max = MaxofInt8(tx.High, tx.Left.Max, tx.Right.Max) //旋转后tx成为ty的子节点，先维护子节点max
	ty.Max = MaxofInt8(ty.High, ty.Left.Max, ty.Right.Max) //再维护父节点max ，右旋类似
}

//右旋
func (itree *ITree) RightRotate(tx *INode) {
	ty := tx.Left
	tx.Left = ty.Right
	if ty.Right != itree.Nil {
		ty.Right.Parent = tx
	}
	ty.Parent = tx.Parent
	if tx.Parent == itree.Nil {
		itree.Root = ty
	} else if tx == tx.Parent.Left {
		tx.Parent.Left = ty
	} else {
		tx.Parent.Right = ty
	}
	ty.Right = tx
	tx.Parent = ty
	tx.Max = MaxofInt8(tx.High, tx.Left.Max, tx.Right.Max)
	ty.Max = MaxofInt8(ty.High, ty.Left.Max, ty.Right.Max)
}

//区间树插入
func (itree *ITree) Insert(node *INode) {
	ty := itree.Nil
	tx := itree.Root
	//通过二叉搜索树找出适合node插入的叶子节点
	for tx != itree.Nil { //tx指向要插入的nil节点，ty指向tx的父节点
		ty = tx
		//if ty.Max < node.Max {
		//	ty.Max = node.Max
		//}
		if node.Low < tx.Low {
			tx = tx.Left
		} else {
			tx = tx.Right
		}
	}
	//将node的父结点指针指向ty
	node.Parent = ty
	if ty == itree.Nil { //若ty是根节点，则直接将根节点指向node
		itree.Root = node
	} else if node.Low < ty.Low { //否则根据Low的值将node分配到ty的左孩子或者右孩子
		ty.Left = node
	} else {
		ty.Right = node
	} //将node进一步调整，融入树结构
	node.Left = itree.Nil
	node.Right = itree.Nil
	node.Icolor = RED
	//调用INsertFixup()来维持红黑性质
	itree.InsertFixup(node)
	//插入前，所有节点max值都满足区间树定义，InsertFixup()中左旋右旋都维护了max，现只需考虑新插入节点对其祖先的max的影响
	for ty != itree.Nil && ty.Max <= node.Max {//从新节点的父节点向上比较，将新节点max值传递到所有需要维护的父节点中
		ty.Max = MaxofInt8(ty.High, ty.Left.Max, ty.Right.Max)
		ty = ty.Parent
	}
}

//区间树插入后维持红黑性质
func (itree *ITree) InsertFixup(node *INode) {
	for node.Parent.Icolor == RED {
		if node.Parent == node.Parent.Parent.Left {
			ty := node.Parent.Parent.Right
			if ty.Icolor == RED {
				node.Parent.Icolor = BLACK
				ty.Icolor = BLACK
				node.Parent.Parent.Icolor = RED
				node = node.Parent.Parent
			} else {

				if node == node.Parent.Right {
					node = node.Parent
					itree.LeftRotate(node)
				}
				node.Parent.Icolor = BLACK
				node.Parent.Parent.Icolor = RED
				itree.RightRotate(node.Parent.Parent)
			}
		} else {
			ty := node.Parent.Parent.Left
			if ty.Icolor == RED {
				node.Parent.Icolor = BLACK
				ty.Icolor = BLACK
				node.Parent.Parent.Icolor = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					itree.RightRotate(node)
				}
				node.Parent.Icolor = BLACK
				node.Parent.Parent.Icolor = RED
				itree.LeftRotate(node.Parent.Parent)
			}
		}
	}
	itree.Root.Icolor = BLACK
}

//子树替换
func (itree *ITree) Transplant(tu, tv *INode) {
	if tu.Parent == itree.Nil {
		itree.Root = tv
	} else if tu == tu.Parent.Left {
		tu.Parent.Left = tv
	} else {
		tu.Parent.Right = tv
	}
	tv.Parent = tu.Parent
}

//区间树删除
func (itree *ITree) Delete(node *INode) {
	tx, ty := itree.Nil, itree.Nil
	ty = node
	ty_original_color := ty.Icolor
	if node.Left == itree.Nil {
		tx = node.Right
		itree.Transplant(node, node.Right)
	} else if node.Right == itree.Nil {
		tx = node.Left
		itree.Transplant(node, node.Left)
	} else {
		ty = itree.Minmum(node.Right)
		ty_original_color = ty.Icolor
		tx = ty.Right
		if ty.Parent == node {
			tx.Parent = ty
		} else {
			itree.Transplant(ty, ty.Right)
			ty.Right = node.Right
			ty.Right.Parent = ty
			//当node左右子树非空且node的后继ty不是node的右子树时，需要将node的后继ty移出原位置，用ty的右子树替换ty
			tg := ty.Parent //此时重新计算ty的父节点的max
			tg.Max = MaxofInt8(tg.High, tg.Right.Max, tg.Left.Max)
		}
		//node左右子树非空，ty为node的后继，用ty替换node，此时ty指向原来node的位置
		itree.Transplant(node, ty)
		ty.Left = node.Left //将node的左孩子与ty建立关系
		ty.Left.Parent = ty
		ty.Icolor = node.Icolor //重新维护新移入节点ty的max
		ty.Max = MaxofInt8(ty.High, ty.Left.Max, ty.Right.Max)
	}
	tg := ty
	//无论要删除节点的子树情况如何，都需要在删除结束后维护其所有祖先的max值，如果其祖先的max与要删除的节点node相同，则重新维护一遍其祖先节点的max
	for tg != itree.Nil && tg.Parent.Max == node.Max{//直到找到某个祖先的max不与要删除节点相同，则该祖先节点以及该祖先节点的祖先的max值都是满足区间树要求的，推出循环即可
		tg = tg.Parent
		tg.Max = MaxofInt8(tg.High, tg.Left.Max, tg.Right.Max)
	}
	if ty_original_color == BLACK {
		itree.DeleteFixup(tx)
	}

}
func (itree *ITree) DeleteFixup(node *INode) {
	for node != itree.Root && node.Icolor == BLACK {
		if node == node.Parent.Left {
			tw := node.Parent.Right
			if tw.Icolor == RED {
				tw.Icolor = BLACK
				node.Parent.Icolor = RED
				itree.LeftRotate(node.Parent)
				tw = node.Parent.Right
			}
			if tw.Left.Icolor == BLACK && tw.Right.Icolor == BLACK {
				tw.Icolor = RED
				node = node.Parent
			} else {
				if tw.Right.Icolor == BLACK {
					tw.Left.Icolor = BLACK
					tw.Icolor = RED
					itree.RightRotate(tw)
					tw = node.Parent.Right
				}
				tw.Icolor = node.Parent.Icolor
				node.Parent.Icolor = BLACK
				tw.Right.Icolor = BLACK
				itree.LeftRotate(node.Parent)
				node = itree.Root
			}
		} else {
			tw := node.Parent.Left
			if tw.Icolor == RED {
				tw.Icolor = BLACK
				node.Parent.Icolor = RED
				itree.RightRotate(node.Parent)
				tw = node.Parent.Left
			}
			if tw.Right.Icolor == BLACK && tw.Left.Icolor == BLACK {
				tw.Icolor = RED
				node = node.Parent
			} else {
				if tw.Left.Icolor == BLACK {
					tw.Right.Icolor = BLACK
					tw.Icolor = RED
					itree.LeftRotate(tw)
					tw = node.Parent.Left
				}
				tw.Icolor = node.Parent.Icolor
				node.Parent.Icolor = BLACK
				tw.Right.Icolor = BLACK
				itree.RightRotate(node.Parent)
				node = itree.Root
			}
		}
	}
	node.Icolor = BLACK
}

//返回是否重叠，若重叠返回true，否则返回false
func (node INode) IsOverlap(l, h int8) bool {
	return l<=node.High && h >= node.Low
}

//区间树查找
func (itree ITree) Search(low, high int8) *INode {
	tx := itree.Root
	for tx != itree.Nil && !tx.IsOverlap(low, high) {
		if tx.Left != itree.Nil && tx.Left.Max >= low {
			tx = tx.Left
		} else {
			tx = tx.Right
		}
	}
	return tx
}

func (itree ITree) LayerTraver() {
	var queue []*INode = make([]*INode, 1024)
	var head, tail int8
	head, tail = 0, 0

	if itree.Root != itree.Nil {
		queue[tail] = itree.Root
		tail += 1
	}
	tmpnode := itree.Root
	for head < tail {
		tmpnode = queue[head]
		if tmpnode.Left != itree.Nil {
			queue[tail] = tmpnode.Left
			tail += 1
		}
		if tmpnode.Right != itree.Nil {
			queue[tail] = tmpnode.Right
			tail += 1
		}
		head += 1
	}
	fmt.Println("Cid,\t Cname,\t\t [Low,\tHigh]\tMax,\tColor,\t Child")
	var i int8
	i = 0
	for i < tail {
		tmpnode := queue[i]
		tmpnode.PrintNode(itree)
		i += 1
	}
}
