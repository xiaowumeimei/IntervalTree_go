package main

import (
	"intervaltree"
)

func main() {
	inode_1 := intervaltree.NewINode(1,2,4,4,"cid = 1", true)
	inode_2 := intervaltree.NewINode(2,0,6,6,"cid = 2", true)
	inode_3 := intervaltree.NewINode(3,1,3,3,"cid = 3", true)
	inode_4 := intervaltree.NewINode(4,4,20,20,"cid = 4", true)
	inode_5 := intervaltree.NewINode(5,9,12,12,"cid = 5", true)
	inode_6 := intervaltree.NewINode(6,6,7,7,"cid = 6", true)
	inode_7 := intervaltree.NewINode(7,12,15,15,"cid = 7", true)
	inode_8 := intervaltree.NewINode(8,3,6,6,"cid = 8", true)
	inode_9 := intervaltree.NewINode(9,7,13,13,"cid = 9", true)
	inode_10 := intervaltree.NewINode(10,10,11,11,"cid = 10", true)
	inode_11 := intervaltree.NewINode(11,3,9,9,"cid = 11", true)
	inode_12 := intervaltree.NewINode(12,2,7,7,"cid = 12", true)
	tree := intervaltree.NewITree()
	tree.Insert(inode_1)
	tree.Insert(inode_2)
	tree.Insert(inode_3)
	tree.Insert(inode_4)
	tree.Insert(inode_5)
	tree.Insert(inode_6)
	tree.Insert(inode_7)
	tree.Insert(inode_8)
	tree.Insert(inode_9)
	tree.Insert(inode_10)
	tree.Insert(inode_11)
	tree.Insert(inode_12)
	tree.Delete(inode_4)
	tree.LayerTraver()
	//tree.Delete(inode_4)
	//tree.LayerTraver()
	//tree.Search(19,20).PrintNode(tree)


	//tree.Delete(inode_4)
	//fmt.Println("\n\nAfter  is \n\n ")
	
}