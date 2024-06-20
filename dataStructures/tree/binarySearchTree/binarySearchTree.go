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

func (b *BinarySearchTree) Find(value any) any {
	return b.Root.find(value).Value
}

func (b *BinarySearchTree) FindMin(value any) any {
	return b.Root.findMin().Value
}

func (b *BinarySearchTree) FindMax(value any) any {
	return b.Root.findMax().Value
}

func (b *BinarySearchTree) Contains(value any) bool {
	return b.Root.Contains(value)
}

func (b *BinarySearchTree) Remove(value any) bool {
	return b.Root.Remove(value)
}

func (b BinarySearchTree) ToString() string {
	return b.Root.ToString()
}

func (b *BinarySearchTree) PrintTree() string {
	t, err := json.MarshalIndent(b, "", "--")

	if err != nil {
		return err.Error()
	}

	return string(t)
}
