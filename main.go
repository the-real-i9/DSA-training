package main

import (
	dll "dst/doublylinkedlist"
	"fmt"
)

func main() {
	// singlylinkedlist.Init()
	dll.Init()

	bh := dll.BrowserHistory{}

	bh.Push(dll.WebPage{Title: "Testing", Body: "Testing this browser history"})
	bh.Push(dll.WebPage{Title: "Testing 22", Body: "Testing another browser history"})

	bh.Back()
	bh.Forward()

	fmt.Printf("%+v\n", bh.Current.Page)
}
