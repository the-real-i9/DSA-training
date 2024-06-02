package tree

import (
	"encoding/json"
	"fmt"
)

type BSTree struct {
	Root   *BTNode
	length int
	cmp    func(x, y any) int
}

func (bst *BSTree) SetCmpFunc(cmpFunc func(x, y any) int) {
	bst.cmp = cmpFunc
}

// the rule of "insert" states: "assign a node if it's empty, else branchInsert"
func (bst *BSTree) Insert(v any) {
	newNode := &BTNode{Data: v}

	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.branchInsert(bst.Root, newNode)
	}

	bst.length++
}

// "branchInsert" states:
// if incoming data is less than that of the parent node, insert to its left child,
// else if incoming data is greater than that of the parent node, insert to its right child
// esle, do nothing, as the data already exists in the tree
func (bst BSTree) branchInsert(dest *BTNode, newNode *BTNode) {

	if bst.cmp(newNode.Data, dest.Data) < 0 {
		if dest.Left == nil {
			dest.Left = newNode
		} else {
			bst.branchInsert(dest.Left, newNode)
		}
	} else if bst.cmp(newNode.Data, dest.Data) > 0 {
		if dest.Right == nil {
			dest.Right = newNode
		} else {
			bst.branchInsert(dest.Right, newNode)
		}
	}
}

func (bst BSTree) Find(target any) (data any, found bool, totalSteps int) {
	if bst.Root == nil {
		return nil, false, 0
	}

	if bst.cmp(target, bst.Root.Data) == 0 {
		return bst.Root.Data, true, 0
	}

	return bst.branchFind(bst.Root, target, 1)
}

func (bst BSTree) branchFind(node *BTNode, target any, steps int) (data any, found bool, totalSteps int) {

	if bst.cmp(target, node.Data) < 0 {
		if node.Left == nil {
			return nil, false, steps
		}

		if bst.cmp(target, node.Left.Data) == 0 {
			return node.Left.Data, true, steps
		}

		return bst.branchFind(node.Left, target, steps+1)
	} else if bst.cmp(target, node.Data) > 0 {
		if node.Right == nil {
			return nil, false, steps
		}

		if bst.cmp(target, node.Right.Data) == 0 {
			return node.Right.Data, true, steps
		}

		return bst.branchFind(node.Right, target, steps+1)
	}

	return nil, false, steps
}

func (bst BSTree) Length() any {
	return bst.length
}

func BSTreeInit() {
	userTree := BSTree{}

	userTree.SetCmpFunc(func(x, y any) int {
		if x.(User).Id < y.(User).Id {
			return -1
		} else if x.(User).Id > y.(User).Id {
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
	userTree.Insert(User{Id: 7052206390, Username: "stu", Name: "Klm Nop"})

	/* res, _ := json.MarshalIndent(userTree, "", "  ")

	fmt.Println(string(res))
	fmt.Println(userTree.Length()) */

	resUser, found, totalSteps := userTree.Find(User{Id: 7052206389})
	if !found {
		fmt.Println("tree.Find: not found")
		return
	}

	data := struct {
		User  any
		Steps int
	}{User: resUser, Steps: totalSteps}

	jsres, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println(string(jsres))

}
