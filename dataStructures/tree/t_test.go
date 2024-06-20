package tree

import (
	"dsa/dataStructures/tree/binarySearchTree"
	"dsa/utils/comparator"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	node := &binarySearchTree.BinarySearchTree{
		NodeValueComparator: comparator.Comparator{
			Compare: comparator.DefaultCompareFunc,
		},
	}

	node.Insert(5)
	node.Insert(6)
	node.Insert(4)

	t.Logf("%+v\n", node.Tree())
}
