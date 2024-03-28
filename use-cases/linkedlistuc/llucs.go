package linkedlistuc

import "fmt"

func BrowserHistoryUC() {
	bh := browserHistory{}

	bh.Push(webPage{Title: "Testing", Body: "Testing this browser history"})
	bh.Push(webPage{Title: "Testing 22", Body: "Testing another browser history"})

	bh.Back()
	bh.Forward()

	fmt.Printf("%+v\n", bh.Current.Page)
}
