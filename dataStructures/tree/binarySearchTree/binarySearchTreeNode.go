package binarySearchTree

import (
	"dsa/utils/comparator"
)

func NewBinarySearchTreeNode(value any, compareFunction func(a, b any) int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		Value:               value,
		compareFunction:     compareFunction,
		nodeComparator:      comparator.NewComparator(nil),
		nodeValueComparator: comparator.NewComparator(compareFunction),
	}
}

type BinarySearchTreeNode struct {
	Value               any
	Left                *BinarySearchTreeNode
	Right               *BinarySearchTreeNode
	parent              *BinarySearchTreeNode
	compareFunction     func(a, b any) int
	nodeComparator      comparator.Comparator
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

		newNode := NewBinarySearchTreeNode(value, b.compareFunction)
		b.SetLeft(newNode)

		return
	}

	if b.nodeValueComparator.GreaterThan(value, b.Value) {
		// Insert to the Right

		if b.Right != nil {
			b.Right.Insert(value)
			return
		}

		newNode := NewBinarySearchTreeNode(value, b.compareFunction)
		b.SetRight(newNode)

	}
}

func (b *BinarySearchTreeNode) Find(value any) any {
	if b.nodeValueComparator.Equal(b.Value, value) {
		return b.Value
	}

	if b.nodeValueComparator.LessThan(value, b.Value) {
		return b.Left.Find(value)
	}

	if b.nodeValueComparator.GreaterThan(value, b.Value) {
		return b.Right.Find(value)
	}

	return nil
}

func (b *BinarySearchTreeNode) Contains(value any) bool {
	return b.Find(value) != nil
}
