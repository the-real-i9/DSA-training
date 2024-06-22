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

var nodeValueCompareFunction = func(a, b any) int {
	if a.(User).Id < b.(User).Id {
		return -1
	} else if a.(User).Id > b.(User).Id {
		return 1
	}

	return 0
}

func TestBinarySearchTree(t *testing.T) {

	bstree := binarySearchTree.NewBinarySearchTree(nil)

	bstree.Insert(User{Id: 8035700462, Username: "i9x", Name: "Kehinde Ogunrinola"})
	bstree.Insert(User{Id: 8056679557, Username: "dollyp", Name: "Dolapo Olaleye"})
	bstree.Insert(User{Id: 8052364849, Username: "menu", Name: "Menu Menu"})
	bstree.Insert(User{Id: 9061928474, Username: "abc", Name: "Abc Def"})
	bstree.Insert(User{Id: 7032639620, Username: "def", Name: "Ghi Jkl"})
	bstree.Insert(User{Id: 8130302370, Username: "ghi", Name: "Mno Pqr"})
	bstree.Insert(User{Id: 8106525254, Username: "jkl", Name: "Stu Vwx"})
	bstree.Insert(User{Id: 8065329258, Username: "mno", Name: "Yza Bcd"})
	bstree.Insert(User{Id: 7052206389, Username: "pqr", Name: "Efg Hij"})
	bstree.Insert(User{Id: 7052206390, Username: "stu", Name: "Klm Nop"})

	bstree.Remove(User{Id: 8065329258})

	/* bstree.Insert(3)
	bstree.Insert(1)
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(15)
	bstree.Insert(13)
	bstree.Insert(12)
	bstree.Insert(9)

	t.Log(bstree.ToString()) */

	// t.Logf("%+v\n", bstree.Contains(User{Id: 8035700462}))
	// t.Log(bstree.PrintTree())
}
