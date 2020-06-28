// Copyright © 2020 JianHui Ding
// Go programming Tree

package main

import (
	"fmt"
	"tree/btree"
)

func main()  {

	{
		//+++++++++++++++++++++++++++++++++++++++++++++++//
		// 普通二叉树练习
		//                 10
		//                /  \
		//               /    \
		//              6      4
		//             / \    /
		//            1   5  2

		// 创建二叉树 思路：先创建节点，然后把节点连接起来
		root := btree.NewBtree(10) // 创建根节点10
		root.Left(6).Right(4).Print()  // 创建一个二叉树 根10 左6 右4
		{
			// 以根节点的左节点为根节点创建一个二叉树  根6 左1 右5
			root.LeftNode.Left(1).Right(5).Print()
			// 以根节点的右节点为根节点创建一个二叉树 根4 左2 右空
			root.RightNode.Left(2).Print()
		}

		// 先序遍历 节点顺序为：根-左-右
		root.PreOrder() // result : ————> 10 ————> 6 ————> 1 ————> 5 ————> 4 ————> 2
		fmt.Print("\n")
		// 中序遍历 节点顺序为：左-根-右
		root.MiddelOrder() // result : ————> 1 ————> 6 ————> 5 ————> 10 ————> 2 ————> 4
		fmt.Print("\n")
		// 后序遍历 节点顺序为：左-右-根
		root.PostOrder() // result : ————> 1 ————> 5 ————> 6 ————> 2 ————> 4 ————> 10
		fmt.Print("\n")

		// 先序遍历的方式计算二叉树的深度
		fmt.Println(root.DepthTree())

		// 画出二叉树
		// 思路：
		//                 10              节点行
		//                /  \             连接线行
		//               /    \
		//              6      4
		//             / \    /
		//            1   5  2
		//           2的幂方
		// 节点行 ： 计算左右空格数：空格————>节点————>空格————>节点————>空格
		//          公式： 左空格数  2(总层数-当前层数)-1  math.Pow(2.0,float64(floor)) - 1
		//          公式： 右空格数  2(总层数-当前层数+1)-1  math.Pow(2.0,float64(floor + 1)) - 1
		// 连接线行: 1、先计算层与层之前的连接线的行数
		//          2、计算左右空格数：空格————>连接线————>空格————>连接线————>空格
		//          公式： 总行数         2(总层数-当前层数-1)          math.Pow(2.0,float64(floor-1))
		//          公式： 左空格数       2(总层数-当前层数)-1-当前行  math.Pow(2.0,float64(floor)) - 1 - i
		//          公式： 连接线间的空格  2*当前行-1             2*i - 1
		//          公式： 右空格数       2*(总行数)-当前行           2*(总行数)-i
		//          公式： 节点为空时 需要打印的空格数  总行数*2+当前行+1
		root.PrintTree(btree.BTrees{root},root.DepthTree(),1)
	}

	{
		//+++++++++++++++++++++++++++++++++++++++++++++++//
		// 二叉查找树(binary search tree) 练习
		// 特征：
		// 　	二叉查找树，也称二叉搜索树，或二叉排序树。 要么是一颗空树，要么就是具有如下性质的二叉树：
		//		（1） 若任意节点的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
		//		（2） 若任意节点的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
		//		（3） 任意节点的左、右子树也分别为二叉查找树；
		//		（4） 没有值相等的节点

		// 创建二叉查找树
		root := btree.NewBSTree(8)
		btree.AddNode(btree.NewBSTree(3),root)
		btree.AddNode(btree.NewBSTree(6),root)
		btree.AddNode(btree.NewBSTree(7),root)
		btree.AddNode(btree.NewBSTree(1),root)
		btree.AddNode(btree.NewBSTree(4),root)
		btree.AddNode(btree.NewBSTree(2),root)
		btree.AddNode(btree.NewBSTree(5),root)
		btree.AddNode(btree.NewBSTree(9),root)
		btree.AddNode(btree.NewBSTree(10),root)
		btree.AddNode(btree.NewBSTree(11),root)
		btree.AddNode(btree.NewBSTree(12),root)


		fmt.Println("//--------------遍历二叉查找树------------------//")
		// 先序遍历 根-左-右
		ret := make([]int,0)
		root.PreOrder(&ret)

		// 中序遍历 左-根-右
		ret1 := make([]int,0)
		root.MiddleOrder(&ret1)

		// 后序遍历 左-右-根
		ret2 := make([]int,0)
		root.PostOrder(&ret2)

		fmt.Println(ret,ret1,ret2)

		// 查找二叉树最大值
		fmt.Println("//--------------二叉查找树最大值------------------//")
		fmt.Println(root.MaxNode())

		// 查找二叉树最小值
		fmt.Println("//--------------二叉查找树最小值------------------//")
		fmt.Println(root.MinNode())

		// 二叉树的深度
		fmt.Println("//--------------二叉查找树的深度------------------//")
		fmt.Println(root.DepthTree())

		// 搜索节点
		fmt.Println("//--------------搜索节点------------------//")
		tree1 := btree.SearchNode(6, root)
		tree1.PrintTree(btree.BSTrees{tree1},tree1.DepthTree(),1)
		tree2 := btree.SearchNode(3, root)
		tree2.PrintTree(btree.BSTrees{tree2},tree2.DepthTree(),1)

		// 搜索父节点
		fmt.Println("//---------------搜索父节点-----------------//")
		_,tree3,_ := btree.SearchNodeWithParent(2, root)
		tree3.PrintTree(btree.BSTrees{tree3},tree3.DepthTree(),1)
		_,tree4,_ := btree.SearchNodeWithParent(4, root)
		tree4.PrintTree(btree.BSTrees{tree4},tree4.DepthTree(),1)


		// 画出二叉查找树
		fmt.Println("//--------------画二叉查找树------------------//")
		root.PrintTree(btree.BSTrees{root},root.DepthTree(),1)


		// 删除二叉查找树节点-没有子节点
		//btree.DelNode(7,root)
		//fmt.Println("//--------------删除节点后------------------//")
		//root.PrintTree(btree.BSTrees{root},root.DepthTree(),1)

		// 删除二叉查找树节点-有一个子节点
		//btree.DelNode(10,root)
		//btree.DelNode(5,root)
		//fmt.Println("//--------------删除节点后------------------//")
		//root.PrintTree(btree.BSTrees{root},root.DepthTree(),1)

		// 删除二叉查找树节点-有两个子节点
		//btree.DelNode(3,root)
		//fmt.Println("//--------------删除节点后------------------//")
		//root.PrintTree(btree.BSTrees{root},root.DepthTree(),1)


		// 删除二叉查找树节点-删除根节点
		//root1 := btree.NewBSTree(8)
		//btree.AddNode(btree.NewBSTree(3),root1)
		//btree.AddNode(btree.NewBSTree(9),root1)

		//btree.DelNode(8,root1)
		//fmt.Println("//--------------删除节点后------------------//")
		//root1.PrintTree(btree.BSTrees{root1},root1.DepthTree(),1)

	}






}
