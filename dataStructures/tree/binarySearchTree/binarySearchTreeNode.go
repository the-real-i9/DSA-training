package binarySearchTree

import (
	"dsa/utils/comparator"
)

type BinarySearchTreeNode struct {
	Left                *BinarySearchTreeNode
	Right               *BinarySearchTreeNode
	Parent              *BinarySearchTreeNode
	Value               any
	NodeValueComparator comparator.Comparator
}

func (b *BinarySearchTreeNode) Insert(value any) {
	if b.NodeValueComparator.Equal(b.Value, nil) {
		b.Value = value
		return
	}

	if b.NodeValueComparator.LessThan(value, b.Value) {
		// Insert to the left

		if b.Left != nil {
			b.Left.Insert(value)
			return
		}

		newNode := &BinarySearchTreeNode{Value: value, NodeValueComparator: b.NodeValueComparator}
		b.SetLeft(newNode)

		return
	}

	if b.NodeValueComparator.GreaterThan(value, b.Value) {
		// Insert to the right

		if b.Right != nil {
			b.Right.Insert(value)
			return
		}

		newNode := &BinarySearchTreeNode{Value: value, NodeValueComparator: b.NodeValueComparator}
		b.SetRight(newNode)

	}
}
