// Copyright © 2020 JianHui Ding
// Go programming Tree

package btree

import (
	"fmt"
	"math"
)

// 普通二叉树练习

type Btree struct {
	Node int       // 根节点
	LeftNode *Btree   // 左节点 *Btree类型 这样左结点下又是一个tree
	RightNode *Btree  // 右节点
}

// 初始化Btree
func NewBtree(node int) *Btree {
	return &Btree{Node:node}
}

// 打印节点
func (this *Btree) Print()  {
	fmt.Printf("根节点为：%d \n", this.Node)

	var lNode interface{}
	// 当前二叉树的左节点存在
	if this.LeftNode != nil {
		lNode = this.LeftNode.Node
	}

	var rNode interface{}
	// 当前二叉树的右节点存在
	if this.RightNode != nil {
		rNode = this.RightNode.Node
	}

	fmt.Printf("左节点为：%v , 右节点为：%v \n", lNode, rNode)
}

// 创建左节点 treeOrnode 可传入tree或单个节点
func (this *Btree) Left(treeOrnode interface{}) *Btree {

	// 断言是否为*Btree类型，是则传入的是一个tree
	if tree,ok := treeOrnode.(*Btree); ok {
		this.LeftNode = tree // 把tree放入当前二叉树的左节点下
	}

	// 断言是否为int类型，是则传入的是单个节点，需要初始化
	if node,ok := treeOrnode.(int); ok {
		this.LeftNode = NewBtree(node) // 初始化当前二叉树的左节点，做为下一级二叉树的根节点
	}

	return this
}

// 创建右节点 treeOrnode 可传入tree或单个节点
func (this *Btree) Right(treeOrnode interface{}) *Btree {
	// 断言是否为*Btree类型，是则传入的是一个tree
	if tree,ok := treeOrnode.(*Btree); ok {
		this.RightNode = tree //把tree放入当前二叉树的右节点下
	}

	// 断言是否为int类型，是则传入的是单个节点，需要初始化
	if node,ok := treeOrnode.(int); ok {
		this.RightNode = NewBtree(node) // 初始化当前二叉树的右节点，做为下一级二叉树的根节点
	}

	return this
}

// 先序遍历所有节点 递归调用 (根-左-右)
func (this *Btree) PreOrder()  {
	// 当前节点为nil时 直接返回
	if this == nil {
		return
	}

	// 当前节点不为nil  打印当前节点
	fmt.Printf("————> %v ", this.Node)

	// 遍历当前节点的左节点下的所有节点
	this.LeftNode.PreOrder()

	// 遍历当前节点的右节点下的所有节点
	this.RightNode.PreOrder()

}

// 中序遍历所有节点 递归调用 (左-根-右)
func (this *Btree) MiddelOrder()  {
	// 当前节点为nil时 直接返回
	if this == nil {
		return
	}

	this.LeftNode.MiddelOrder()

	fmt.Printf("————> %v ", this.Node)

	this.RightNode.MiddelOrder()

}

// 后序遍历所有节点 递归调用 (左-右-根)
func (this *Btree) PostOrder()  {
	// 当前节点为nil时 直接返回
	if this == nil {
		return
	}

	this.LeftNode.PostOrder()

	this.RightNode.PostOrder()

	fmt.Printf("————> %v ", this.Node)

}

// 计算二叉树的高度，递归
func (this *Btree) DepthTree() int {
	// 当前节点为nil时 返回0 高度
	if this == nil {
		return 0
	}

	//  计算当前节点下的左节点深度
	lDepth := this.LeftNode.DepthTree()
	//  计算当前节点下的右节点深度
	rDepth := this.RightNode.DepthTree()

	// 左右节点的深度对比 返回大者 并 + 1
	if lDepth >= rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

// 打印空格
func (this *Btree) PrintBlanks(count float64)  {
	for i := 1.0; i <= count; i++ {
		fmt.Print(" ")
	}
}

// 画二叉树   trees 二叉树集合  maxLevel 二叉树的深度(层高) currLevel 当前层级
func (this *Btree) PrintTree(trees BTrees, maxLevel int, currLevel int)  {
	// tree为空或没有子节点了，直接返回
	if len(trees) == 0 || trees.IsAllNil() {
		return
	}

	floor := maxLevel - currLevel
	lblanks := math.Pow(2.0,float64(floor)) - 1  // 节点左边空格数
	rblanks := math.Pow(2.0,float64(floor + 1)) - 1 // 节点右边空格数

	newTrees := make(BTrees, 0) //  新切片 用于递归子节点

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
func (this *Btree) PrintLine(trees BTrees, lineNums float64, blanks float64)  {
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
