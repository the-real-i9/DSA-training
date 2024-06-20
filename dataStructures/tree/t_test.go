package tree

import (
	"dsa/dataStructures/tree/binarySearchTree"
	"testing"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func TestBinarySearchTree(t *testing.T) {

	bstree := binarySearchTree.NewBinarySearchTree(nil)

	bstree.Insert(5)
	bstree.Insert(6)
	bstree.Insert(4)
	bstree.Insert(3)

	t.Logf("%d\n", bstree.Find(3))
}
