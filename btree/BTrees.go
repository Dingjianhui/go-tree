// Copyright © 2020 JianHui Ding
// Go programming Tree

package btree

// 二叉树集合类型
type BTrees []*Btree

func (this BTrees) IsAllNil() bool {
	for _,tree := range this {
		if tree != nil {
			return false
		}
	}
	return true
}

type BSTrees []*BSTree

func (this BSTrees) IsAllNil() bool {
	for _,tree := range this {
		if tree != nil {
			return false
		}
	}
	return true
}

