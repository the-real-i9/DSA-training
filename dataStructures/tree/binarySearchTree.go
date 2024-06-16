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

func (bst *BSTree) Delete(value any) {

}

func (bst BSTree) Search(target any) (data any, found bool, totalSteps int) {
	if bst.Root == nil {
		return nil, false, 0
	}

	if bst.cmp(target, bst.Root.Data) == 0 {
		return bst.Root.Data, true, 0
	}

	return bst.branchSearch(bst.Root, target, 1)
}

func (bst BSTree) branchSearch(node *BTNode, target any, steps int) (data any, found bool, totalSteps int) {

	if bst.cmp(target, node.Data) < 0 {
		if node.Left == nil {
			return nil, false, steps
		}

		if bst.cmp(target, node.Left.Data) == 0 {
			return node.Left.Data, true, steps
		}

		return bst.branchSearch(node.Left, target, steps+1)
	}

	if node.Right == nil {
		return nil, false, steps
	}

	if bst.cmp(target, node.Right.Data) == 0 {
		return node.Right.Data, true, steps
	}

	return bst.branchSearch(node.Right, target, steps+1)

}

func (bst BSTree) Traverse(out chan<- any) {
	if bst.Root == nil {
		return
	}

	out <- bst.Root.Data

	bst.branchTraverse(bst.Root, out)

	close(out)
}

func (bst BSTree) branchTraverse(node *BTNode, out chan<- any) {

	if node.Left != nil {
		out <- node.Left.Data
		bst.branchTraverse(node.Left, out)
	}

	if node.Right != nil {
		out <- node.Right.Data
		bst.branchTraverse(node.Right, out)
	}
}

func (bst BSTree) Min() any {
	recv := make(chan any)
	var min any

	go func() {
		bst.Traverse(recv)
	}()

	for val := range recv {
		if min != nil {
			if bst.cmp(val, min) < 0 {
				min = val
			}
		} else {
			min = val
		}
	}

	return min
}

func (bst BSTree) Max() any {
	recv := make(chan any)
	var max any

	go func() {
		bst.Traverse(recv)
	}()

	for val := range recv {
		if max != nil {
			if bst.cmp(val, max) > 0 {
				max = val
			}
		} else {
			max = val
		}
	}

	return max
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

	// Insert
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

	// Search
	/* resUser, found, totalSteps := userTree.Search(User{Id: 7052206389})
	if !found {
		fmt.Println("tree.Search: not found")
		return
	}

	data := struct {
		User  any
		Steps int
	}{User: resUser, Steps: totalSteps}

	jsres, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println(string(jsres)) */

	// Traverse
	/* userChan := make(chan any)

	go func() {
		userTree.Traverse(userChan)
	}()

	for user := range userChan {

		fmt.Println(user.(User))
		time.Sleep(1 * time.Second)
	} */

	minUser := userTree.Min()
	maxUser := userTree.Max()

	data := struct {
		MinUser any `json:"minUser"`
		MaxUser any `json:"maxUser"`
	}{MinUser: minUser, MaxUser: maxUser}

	jsres, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println(string(jsres))
}
