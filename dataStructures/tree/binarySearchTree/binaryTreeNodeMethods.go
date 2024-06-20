package binarySearchTree

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

// A node's parent's sibling
func (b BinarySearchTreeNode) Uncle() *BinarySearchTreeNode {
	// Check if the current node has parent
	if b.parent == nil {
		return nil
	}

	// Check if current node has grand-parent
	if b.parent.parent == nil {
		return nil
	}

	// check if grand-parent has two children
	if b.parent.parent.Left == nil || b.parent.parent.Right == nil {
		return nil
	}

	// Now that we know that the current node has grand-parent and this grand-parent has two childeren

	// if this node's parent is the grand-parent's Left child,
	if b.nodeComparator.Equal(b.parent, b.parent.parent.Left) {
		//  then the uncle is its Right child,
		return b.parent.parent.Right
	}

	// else the uncle is its Left child
	return b.parent.parent.Left
}

func (b *BinarySearchTreeNode) SetValue(value any) {
	b.Value = value

}

func (b *BinarySearchTreeNode) SetLeft(node *BinarySearchTreeNode) {
	// if node (parent) currently has a Left child
	if b.Left != nil {

		// let the Left child know it no longer has a parent, by
		// removing the Left child's parent reference to this node (parent)
		b.Left.parent = nil
	}

	// set this node's (parent's) new Left child
	b.Left = node

	// let the new Left child know it now has a parent, by
	// setting the new Left child's parent reference to this node (parent)
	if node != nil {
		b.Left.parent = b
	}
}

// just like SetLeft, but for the Right child in this case
func (b *BinarySearchTreeNode) SetRight(node *BinarySearchTreeNode) {
	if b.Right != nil {
		b.Right.parent = nil
	}

	b.Right = node

	if node != nil {
		b.Right.parent = b
	}
}

func (b *BinarySearchTreeNode) RemoveChild(nodeToRemove *BinarySearchTreeNode) bool {
	if b.Left != nil && b.nodeComparator.Equal(b.Left, nodeToRemove) {
		b.Left = nil
		return true
	}

	if b.Right != nil && b.nodeComparator.Equal(b.Right, nodeToRemove) {
		b.Right = nil
		return true
	}

	return false
}

func (b *BinarySearchTreeNode) ReplaceChild(nodeToReplace *BinarySearchTreeNode, replacementNode *BinarySearchTreeNode) bool {
	if nodeToReplace == nil || replacementNode == nil {
		return false
	}

	if b.Left != nil && b.nodeComparator.Equal(b.Left, nodeToReplace) {
		b.Left = replacementNode
		return true
	}

	if b.Right != nil && b.nodeComparator.Equal(b.Right, nodeToReplace) {
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

// In-Order traversal starts from Left child, to parent or root, and, finally, to Right child
// It traverses the tree, presumably, as a Balanced Tree
func (b BinarySearchTreeNode) TraverseInOrder() []any {
	res := []any{}

	// Add Left node
	if b.Left != nil {
		res = append(res, b.Left.TraverseInOrder()...)
	}

	// Add parent or root
	res = append(res, b.Value)

	// Add Right
	if b.Right != nil {
		res = append(res, b.Right.TraverseInOrder()...)
	}

	return res
}
