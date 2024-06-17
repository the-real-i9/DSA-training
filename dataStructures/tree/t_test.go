package tree

import (
	"testing"
)

func TestBTNode(t *testing.T) {
	node := &BinaryTreeNode{}
	node2 := &BinaryTreeNode{}
	node3 := &BinaryTreeNode{}

	node.SetValue(10)
	node2.SetValue(5)
	node3.SetValue(15)

	node.SetLeft(node2)
	node.SetRight(node3)

	t.Logf("%+v\n", node.String())
}
