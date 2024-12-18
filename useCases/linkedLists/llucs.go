package lluc

import (
	"dsa/useCases/linkedLists/webBrowser"
	"fmt"
)

func UseBrowserHistory() {
	bh := &webBrowser.History{}

	bh.Push(webBrowser.WebPage{Title: "Testing", Body: "Testing this browser history"})
	bh.Push(webBrowser.WebPage{Title: "Testing 22", Body: "Testing another browser history"})

	bh.Back()
	bh.Forward()

	fmt.Printf("%+v\n", bh.Current.Main)
}
