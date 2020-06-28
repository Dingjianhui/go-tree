// Copyright © 2020 JianHui Ding
// Go programming Tree


package btree

import (
	"fmt"
	"math"
)

// 二叉查找树(binary search tree)

type BSTree struct {
	Node int       // 根节点
	LeftNode *BSTree   // 左节点
	RightNode *BSTree  // 右节点
}

func NewBSTree(node int) *BSTree {
	return &BSTree{Node:node}
}

// 动态创建节点
func AddNode(tree *BSTree, root *BSTree) *BSTree {
	if root == nil {
		return tree
	}
	// 根节点  与 新节点对比
	// 新节点比 根节点大 放入右节点
	// 新节点比 根节点小 放入左节点
	if tree.Node > root.Node {
		root.RightNode = AddNode(tree,root.RightNode)
	} else if tree.Node < root.Node {
		root.LeftNode = AddNode(tree,root.LeftNode)
	} else {
		return root
	}

	return root
}

// 搜索结点
func SearchNode(node int, tree *BSTree) *BSTree {
	if tree == nil {
		return nil
	}
	if node < tree.Node {
		return SearchNode(node, tree.LeftNode) // 递归查找左节点
	} else if node > tree.Node {
		return  SearchNode(node, tree.RightNode) // 递归查找右节点
	} else {
		return tree
	}
}

// 搜索父节点
func SearchNodeWithParent(node int,tree *BSTree,parentNode ...interface{}) (*BSTree,*BSTree,string) {
	if tree == nil {
		return nil,nil,""
	}
	if node < tree.Node {
		return SearchNodeWithParent(node,tree.LeftNode,tree,"left")  // 递归查找左节点
	} else if node > tree.Node {
		return SearchNodeWithParent(node,tree.RightNode,tree,"right")  // 递归查找右节点
	} else {
		if len(parentNode) == 0 {
			return tree,nil,""
		} else {
			return tree,parentNode[0].(*BSTree),parentNode[1].(string)
		}
	}
}

// 删除节点 (没有子节点、有一个子节点、有两个子节点、删除根节点)
func DelNode(node int, tree *BSTree)  {
	nodeTree,parentTree,child := SearchNodeWithParent(node, tree)

	if nodeTree == nil {
		return  // 节点没找到
	}

	//if parentTree == nil {
	//	return
	//}

	if nodeTree.IsLeaf() {  // 删除没有子节点的
		// 如果删除的是根节点-没有子节点
		if parentTree == nil {
			*tree = *(*BSTree)(nil) // 先将nil转化为*BSTree类型
			return
		}

		// 如果是左节点，那么左节点置为空
		if child == "left" {
			parentTree.LeftNode = nil
		} else {
			parentTree.RightNode = nil
		}
	} else if single := nodeTree.GetSingleNode(); single != nil {  // 删除有一个子节点的
		// 如果删除的是根节点,只有一个子节点时 , 需要用指针,要不然根节点的值不会被变化
		if parentTree == nil {
			if nodeTree.LeftNode != nil {
				*tree = *nodeTree.LeftNode
			} else {
				*tree = *nodeTree.RightNode
			}
			return
		}



		if child == "left" {
			parentTree.LeftNode = single
		} else {
			parentTree.RightNode = single
		}
	} else {  // 有两个子节点
		value := nodeTree.RightNode.MinNode() // 找到右节点的最小值
		DelNode(value,tree) // 删除右节点的最小值
		nodeTree.Node = value // 要删除的节点的值更改为最小值
	}
}

// 获取子节点
func (this *BSTree) GetSingleNode() *BSTree  {
	if this.LeftNode != nil && this.RightNode == nil {
		return this.LeftNode
	}

	if this.LeftNode == nil && this.RightNode != nil {
		return this.RightNode
	}

	return nil
}

// 判断节点是否有子节点、有没有叶子
func (this *BSTree) IsLeaf() bool {
	// 没有节点
	if this.LeftNode == nil && this.RightNode == nil {
		return true
	} else {
		return false
	}
}

// 先序 根-左-右
func (this *BSTree) PreOrder(ret *[]int)  {
	if this == nil {
		return
	}
	*ret = append(*ret,this.Node)
	this.LeftNode.PreOrder(ret)
	this.RightNode.PreOrder(ret)
}

// 中序 左-根-右
func (this *BSTree) MiddleOrder(ret *[]int)  {
	if this == nil {
		return
	}
	this.LeftNode.MiddleOrder(ret)
	*ret = append(*ret,this.Node)
	this.RightNode.MiddleOrder(ret)
}

// 后序 左-右-根
func (this *BSTree) PostOrder(ret *[]int)  {
	if this == nil {
		return
	}
	this.LeftNode.PostOrder(ret)
	this.RightNode.PostOrder(ret)
	*ret = append(*ret,this.Node)
}

// 二叉树深度
func (this *BSTree) DepthTree() int {
	if this == nil {
		return 0
	}

	lDepth := this.LeftNode.DepthTree()

	rDepth := this.RightNode.DepthTree()

	if lDepth >= rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

// 查找最大值
func (this *BSTree) MaxNode() int {
	if this.RightNode != nil {
		return this.RightNode.MaxNode()
	} else {
		return this.Node
	}
}

// 查找最小值
func (this *BSTree) MinNode() int {
	if this.LeftNode != nil {
		return this.LeftNode.MinNode()
	} else {
		return this.Node
	}
}

// 打印空格
func (this *BSTree) PrintBlanks(count float64)  {
	for i := 1.0; i <= count; i++ {
		fmt.Print(" ")
	}
}

// 画二叉树   trees 二叉树集合  maxLevel 二叉树的深度(层高) currLevel 当前层级
func (this *BSTree) PrintTree(trees BSTrees, maxLevel int, currLevel int)  {
	// tree为空或没有子节点了，直接返回
	if len(trees) == 0 || trees.IsAllNil() {
		return
	}

	floor := maxLevel - currLevel
	lblanks := math.Pow(2.0,float64(floor)) - 1  // 节点左边空格数
	rblanks := math.Pow(2.0,float64(floor + 1)) - 1 // 节点右边空格数

	newTrees := make(BSTrees, 0) //  新切片 用于递归子节点

	this.PrintBlanks(lblanks) // 打印节点左边空格数

	for _,tree := range trees {
		if tree != nil {
			fmt.Print(tree.Node) // 打印节点
			newTrees = append(newTrees, tree.LeftNode, tree.RightNode) // 当前节点的子节点
		} else {
			this.PrintBlanks(1) // 节点为空时打印一个空格数
			newTrees = append(newTrees, nil, nil) // 子节点都为nil时返回
		}
		this.PrintBlanks(rblanks) // 打印右节点
	}

	fmt.Print("\n") // 换行

	// 画线
	lineNums := math.Pow(2.0,float64(floor - 1)) // 画线的行数
	this.PrintLine(trees,lineNums,lblanks)

	this.PrintTree(newTrees, maxLevel, currLevel + 1) // 递归打印当前节点下的子节点
}

// 打印连接线 trees 节点集合  lineNums 线的行数
func (this *BSTree) PrintLine(trees BSTrees, lineNums float64, blanks float64)  {
	for i := 1.0; i <= lineNums; i++ {
		for _,tree := range trees {
			this.PrintBlanks(blanks - i) // 打印连接线左边的空格数
			if tree == nil {
				this.PrintBlanks(lineNums * 2 + i + 1) // 节点为空时需要打印的空格数
				continue
			}
			// 有左节点，那么画左线
			if tree.LeftNode != nil {
				fmt.Print("/")
			} else {
				//fmt.Print("/")
				this.PrintBlanks(1)
			}

			// 线与线之间的空格
			this.PrintBlanks(2*i - 1)

			// 有右节点，那么画右线
			if tree.RightNode != nil {
				fmt.Print("\\")
			} else {
				//fmt.Print("\\")
				this.PrintBlanks(1)
			}

			this.PrintBlanks(2 * lineNums - i) // 打印连接线的右边空格

		}
		// 打印完一行，换行
		fmt.Print("\n")
	}
}
