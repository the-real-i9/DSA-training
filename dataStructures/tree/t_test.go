package tree

import (
	"dsa/dataStructures/tree/binarySearchTree"
	"dsa/utils/comparator"
	"testing"
)

func TestBSTNode(t *testing.T) {

	node := &binarySearchTree.BinarySearchTreeNode{
		NodeValueComparator: comparator.Comparator{
			Compare: comparator.DefaultCompareFunc,
		},
	}

	node.Insert(5)
	node.Insert(6)
	node.Insert(4)
	node.Left.Insert(3)

	t.Logf("%+v\n", node)
}
