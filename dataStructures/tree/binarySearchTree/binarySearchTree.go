package binarySearchTree

import "dsa/utils/comparator"

type BinarySearchTree struct {
	Root                *BinarySearchTreeNode
	NodeValueComparator comparator.Comparator
}

func (b *BinarySearchTree) Insert(value any) {
	newNode := &BinarySearchTreeNode{Value: value, nodeValueComparator: b.NodeValueComparator}

	if b.Root == nil {
		b.Root = newNode
		return
	}

	b.Root.Insert(value)
}

func (b *BinarySearchTree) Tree() string {
	return b.Root.Tree()
}
