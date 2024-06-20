package binarySearchTree

import (
	"dsa/utils/comparator"
)

func NewBinarySearchTreeNode(value any, compareFunction func(a, b any) int) *BinarySearchTreeNode {
	nodeCompareFunction := func(a, b any) int {
		return compareFunction(a.(*BinarySearchTreeNode).Value, b.(*BinarySearchTreeNode).Value)
	}

	return &BinarySearchTreeNode{
		Value:               value,
		compareFunction:     compareFunction,
		nodeComparator:      comparator.NewComparator(nodeCompareFunction),
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

func (b BinarySearchTreeNode) Find(value any) any {
	return b.find(value).Value
}

func (b *BinarySearchTreeNode) find(value any) *BinarySearchTreeNode {
	if b.nodeValueComparator.Equal(b.Value, value) {
		return b
	}

	if b.nodeValueComparator.LessThan(value, b.Value) {
		if b.Left != nil {
			return b.Left.find(value)
		}

		return nil
	}

	if b.nodeValueComparator.GreaterThan(value, b.Value) {
		if b.Right != nil {
			return b.Right.find(value)
		}

		return nil
	}

	return nil
}

func (b *BinarySearchTreeNode) Contains(value any) bool {
	return b.find(value) != nil
}

func (b *BinarySearchTreeNode) Remove(value any) bool {
	nodeToRemove := b.find(value)

	if nodeToRemove == nil {
		return false
	}

	parent := nodeToRemove.parent

	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		// Node is a leaf and thus has no children
		if parent != nil {
			// Node has a parent. Remove this node from the parent
			parent.RemoveChild(nodeToRemove)
		} else {
			// Node has no parent (i.e. root). Just erase the current node value (making it to an empty tree)
			nodeToRemove.SetValue(nil)
		}
	} else if nodeToRemove.Left != nil && nodeToRemove.Right != nil {
		// Node has two children
		// Find the next biggest value (minimum value in the right branch)
		// and replace current value node with that next biggest value
		nextBiggerNode := nodeToRemove.Right.findMin()
		if !b.nodeComparator.Equal(nextBiggerNode, nodeToRemove.Right) {
			b.Remove(nextBiggerNode.Value)
			nodeToRemove.SetValue(nextBiggerNode.Value)
		} else {
			// In case if next right value is the next bigger one and it doesn't have left child
			// then just replace the node that is going to be deleted with the right node
			nodeToRemove.SetValue(nodeToRemove.Right.Value)
			nodeToRemove.SetRight(nodeToRemove.Right.Right)
		}
	} else {
		// Node has only one child
		// Make this child to be a direct child of current node's parent
		childNode := nodeToRemove.Left
		if childNode == nil {
			childNode = nodeToRemove.Right
		}

		if parent != nil {
			parent.ReplaceChild(nodeToRemove, childNode)
		} else {
			CopyNode(nodeToRemove, childNode)
		}
	}

	// Clear the parent of removed node
	nodeToRemove.parent = nil

	return true
}

func (b *BinarySearchTreeNode) findMin() *BinarySearchTreeNode {
	if b.Left == nil {
		return b
	}

	return b.Left.findMin()
}

func (b *BinarySearchTreeNode) findMax() *BinarySearchTreeNode {
	if b.Right == nil {
		return b
	}

	return b.Right.findMax()
}
