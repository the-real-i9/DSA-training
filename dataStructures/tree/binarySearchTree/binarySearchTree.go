package binarySearchTree

import (
	"dsa/utils/comparator"
	"encoding/json"
)

func NewBinarySearchTree(nodeValueCompareFunction func(a, b any) int) *BinarySearchTree {
	root := NewBinarySearchTreeNode(nil, nodeValueCompareFunction)

	return &BinarySearchTree{Root: root, nodeComparator: root.nodeComparator}
}

type BinarySearchTree struct {
	Root           *BinarySearchTreeNode
	nodeComparator comparator.Comparator
}

func (b *BinarySearchTree) Insert(value any) {
	b.Root.Insert(value)
}

func (b *BinarySearchTree) PrintTree() string {
	t, err := json.MarshalIndent(b, "", "--")

	if err != nil {
		return err.Error()
	}

	return string(t)
}
