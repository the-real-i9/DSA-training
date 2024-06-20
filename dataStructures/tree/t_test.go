package tree

import (
	"dsa/dataStructures/tree/binarySearchTree"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	bstree := binarySearchTree.NewBinarySearchTree(nil)

	bstree.Insert(5)
	bstree.Insert(6)
	bstree.Insert(4)

	t.Logf("%s\n", bstree.PrintTree())
}
