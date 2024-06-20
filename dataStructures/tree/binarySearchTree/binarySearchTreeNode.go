package binarySearchTree

import (
	"dsa/utils/comparator"
)

type BinarySearchTreeNode struct {
	Value               any
	Left                *BinarySearchTreeNode
	Right               *BinarySearchTreeNode
	parent              *BinarySearchTreeNode
	nodeValueComparator comparator.Comparator
}

func (b *BinarySearchTreeNode) Insert(value any) {
	if b.nodeValueComparator.Equal(b.Value, nil) {
		b.Value = value
		return
	}

	if b.nodeValueComparator.LessThan(value, b.Value) {
		// Insert to the Left

		if b.Left != nil {
			b.Left.Insert(value)
			return
		}

		newNode := &BinarySearchTreeNode{Value: value, nodeValueComparator: b.nodeValueComparator}
		b.SetLeft(newNode)

		return
	}

	if b.nodeValueComparator.GreaterThan(value, b.Value) {
		// Insert to the Right

		if b.Right != nil {
			b.Right.Insert(value)
			return
		}

		newNode := &BinarySearchTreeNode{Value: value, nodeValueComparator: b.nodeValueComparator}
		b.SetRight(newNode)

	}
}
