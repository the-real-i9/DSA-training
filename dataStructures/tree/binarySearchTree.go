package tree

import (
	"encoding/json"
	"fmt"
)

type BSTree struct {
	Root   *BTNode
	length int
	cmp    func(incValue, nodeValue any) int
}

func (bst *BSTree) SetCmpFunc(cmpFunc func(incValue, nodeValue any) int) {
	bst.cmp = cmpFunc
}

func (bst *BSTree) Insert(v any) {
	newNode := &BTNode{Data: v}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.branchInsert(bst.Root, v)
	}

	bst.length++
}

func (bst BSTree) branchInsert(dest *BTNode, v any) {
	newNode := &BTNode{Data: v}

	if bst.cmp(v, dest.Data) < 0 {
		if dest.Left == nil {
			dest.Left = newNode
		} else {
			bst.branchInsert(dest.Left, v)
		}
	} else if bst.cmp(v, dest.Data) > 0 {
		if dest.Right == nil {
			dest.Right = newNode
		} else {
			bst.branchInsert(dest.Right, v)
		}
	}
}

func (bst *BSTree) Length() any {
	return bst.length
}

func BSTreeInit() {
	userTree := BSTree{}

	userTree.SetCmpFunc(func(incValue, nodeValue any) int {
		if incValue.(User).Id < nodeValue.(User).Id {
			return -1
		} else if incValue.(User).Id > nodeValue.(User).Id {
			return 1
		}

		return 0
	})

	userTree.Insert(User{Id: 8035700462, Username: "i9x", Name: "Kehinde Ogunrinola"})
	userTree.Insert(User{Id: 8056679557, Username: "dollyp", Name: "Dolapo Olaleye"})
	userTree.Insert(User{Id: 8052364849, Username: "menu", Name: "Menu Menu"})
	userTree.Insert(User{Id: 9061928474, Username: "abc", Name: "Abc Def"})
	userTree.Insert(User{Id: 7032639620, Username: "def", Name: "Ghi Jkl"})
	userTree.Insert(User{Id: 8130302370, Username: "ghi", Name: "Mno Pqr"})
	userTree.Insert(User{Id: 8106525254, Username: "jkl", Name: "Stu Vwx"})
	userTree.Insert(User{Id: 8065329258, Username: "mno", Name: "Yza Bcd"})
	userTree.Insert(User{Id: 7052206389, Username: "pqr", Name: "Efg Hij"})

	res, _ := json.MarshalIndent(userTree, "", "  ")

	fmt.Println(string(res))
	fmt.Println(userTree.Length())
}
