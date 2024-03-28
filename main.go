package main

import (
	"dst/doublylinkedlist"
	"usecases/linkedlistuc"
)

func main() {
	// singlylinkedlist.Init()
	doublylinkedlist.Init()
	linkedlistuc.BrowserHistoryUC()
}
