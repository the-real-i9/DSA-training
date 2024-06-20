package binarySearchTree

import (
	"fmt"
)

/* type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
} */

func (b BinarySearchTreeNode) LeftHeight() int {
	if b.Left == nil {
		return 0
	}

	return b.Left.Height()
}

func (b BinarySearchTreeNode) RightHeight() int {
	if b.Right == nil {
		return 0
	}

	return b.Right.Height()
}

func (b BinarySearchTreeNode) Height() int {
	return max(b.LeftHeight(), b.RightHeight()) + 1
}

func (b BinarySearchTreeNode) BalanceFactor() int {
	return b.LeftHeight() - b.RightHeight()
}

// A node's Parent's sibling
func (b BinarySearchTreeNode) Uncle() *BinarySearchTreeNode {
	// Check if the current node has Parent
	if b.Parent == nil {
		return nil
	}

	// Check if current node has grand-Parent
	if b.Parent.Parent == nil {
		return nil
	}

	// check if grand-Parent has two children
	if b.Parent.Parent.Left == nil || b.Parent.Parent.Right == nil {
		return nil
	}

	// Now that we know that the current node has grand-Parent and this grand-Parent has two childeren

	// if this node's Parent is the grand-Parent's Left child,
	if b.NodeValueComparator.Equal(b.Parent, b.Parent.Parent.Left) {
		//  then the uncle is its Right child,
		return b.Parent.Parent.Right
	}

	// else the uncle is its Left child
	return b.Parent.Parent.Left
}

func (b *BinarySearchTreeNode) SetValue(value any) {
	b.Value = value

}

func (b *BinarySearchTreeNode) SetLeft(node *BinarySearchTreeNode) {
	// if node (Parent) currently has a Left child
	if b.Left != nil {

		// let the Left child know it no longer has a Parent, by
		// removing the Left child's Parent reference to this node (Parent)
		b.Left.Parent = nil
	}

	// set this node's (Parent's) new Left child
	b.Left = node

	// let the new Left child know it now has a Parent, by
	// setting the new Left child's Parent reference to this node (Parent)
	if node != nil {
		b.Left.Parent = b
	}
}

// just like SetLeft, but for the Right child in this case
func (b *BinarySearchTreeNode) SetRight(node *BinarySearchTreeNode) {
	if b.Right != nil {
		b.Right.Parent = nil
	}

	b.Right = node

	if node != nil {
		b.Right.Parent = b
	}
}

func (b *BinarySearchTreeNode) RemoveChild(nodeToRemove *BinarySearchTreeNode) bool {
	if b.Left != nil && b.NodeValueComparator.Equal(b.Left, nodeToRemove) {
		b.Left = nil
		return true
	}

	if b.Right != nil && b.NodeValueComparator.Equal(b.Right, nodeToRemove) {
		b.Right = nil
		return true
	}

	return false
}

func (b *BinarySearchTreeNode) ReplaceChild(nodeToReplace *BinarySearchTreeNode, replacementNode *BinarySearchTreeNode) bool {
	if nodeToReplace == nil || replacementNode == nil {
		return false
	}

	if b.Left != nil && b.NodeValueComparator.Equal(b.Left, nodeToReplace) {
		b.Left = replacementNode
		return true
	}

	if b.Right != nil && b.NodeValueComparator.Equal(b.Right, nodeToReplace) {
		b.Right = replacementNode
		return true
	}

	return false
}

func CopyNode(targetNode *BinarySearchTreeNode, sourceNode BinarySearchTreeNode) {
	targetNode.SetValue(sourceNode.Value)
	targetNode.SetLeft(targetNode.Left)
	targetNode.SetRight(targetNode.Right)
}

// In-Order traversal starts from Left child, to Parent or root, and, finally, to Right child
// It traverses the tree, presumably, as a Balanced Tree
func (b BinarySearchTreeNode) TraverseInOrder() []any {
	res := []any{}

	// Add Left node
	if b.Left != nil {
		res = append(res, b.Left.TraverseInOrder()...)
	}

	// Add Parent or root
	res = append(res, b.Value)

	// Add Right
	if b.Right != nil {
		res = append(res, b.Right.TraverseInOrder()...)
	}

	return res
}

func (b BinarySearchTreeNode) ToString() string {
	return fmt.Sprint(b.TraverseInOrder())
}
