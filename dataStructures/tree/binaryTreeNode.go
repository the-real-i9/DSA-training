package tree

import "dsa/utils/comparator"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type BinaryTreeNode struct {
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
	Value  any

	// Any node related meta information may be stored here
	Meta map[string]any

	// This comparator is used to compare binary tree nodes with each other
	NodeComparator comparator.Comparator
}

func (b BinaryTreeNode) LeftHeight() int {
	if b.Left == nil {
		return 0
	}

	return b.Left.Height() + 1
}

func (b BinaryTreeNode) RightHeight() int {
	if b.Left == nil {
		return 0
	}

	return b.Right.Height() + 1
}

func (b BinaryTreeNode) Height() int {
	return max(b.LeftHeight(), b.RightHeight())
}

func (b BinaryTreeNode) BalanceFactor() int {
	return b.LeftHeight() - b.RightHeight()
}

// A node's parent's sibling
func (b BinaryTreeNode) Uncle() *BinaryTreeNode {
	// Check if the current node has parent
	if b.Parent == nil {
		return nil
	}

	// Check if current node has grand-parent
	if b.Parent.Parent == nil {
		return nil
	}

	// check if grand-parent has two children
	if b.Parent.Parent.Left == nil || b.Parent.Parent.Right == nil {
		return nil
	}

	// Now that we know that the current node has grand-parent and this grand-parent has two childeren

	// if this node's parent is the grand-parent's left child,
	if b.NodeComparator.Equal(b.Parent, b.Parent.Parent.Left) {
		//  then the uncle is its right child,
		return b.Parent.Parent.Right
	}

	// else the uncle is its left child
	return b.Parent.Parent.Left
}

func (b *BinaryTreeNode) SetValue(value any) {
	b.Value = value

}

func (b *BinaryTreeNode) SetLeft(node *BinaryTreeNode) {
	// if node (parent) currently has a left child
	if b.Left != nil {

		// let the left child know it no longer has a parent, by
		// removing the left child's parent reference to this node (parent)
		b.Left.Parent = nil
	}

	// set this node's (parent's) new left child
	b.Left = node

	// let the new left child know it now has a parent, by
	// setting the new left child's parent reference to this node (parent)
	if node != nil {
		b.Left.Parent = b
	}
}

// just like SetLeft, but for the right child in this case
func (b *BinaryTreeNode) SetRight(node *BinaryTreeNode) {
	if b.Right != nil {
		b.Right.Parent = nil
	}

	b.Right = node

	if node != nil {
		b.Right.Parent = b
	}
}

func (b *BinaryTreeNode) RemoveChild(nodeToRemove BinaryTreeNode) bool {
	if b.Left != nil && b.NodeComparator.Equal(b.Left, nodeToRemove) {
		b.Left = nil
		return true
	}

	if b.Right != nil && b.NodeComparator.Equal(b.Right, nodeToRemove) {
		b.Right = nil
		return true
	}

	return false
}
