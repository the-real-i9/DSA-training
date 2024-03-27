package doublylinkedlist

type WebPage struct {
	Title string
	Body  string
}

type browserPage struct {
	prev *browserPage
	Page WebPage
	next *browserPage
}

type BrowserHistory struct {
	head    *browserPage
	tail    *browserPage
	Current *browserPage
	length  int
}

func (bh *BrowserHistory) Push(page WebPage) {
	newPage := &browserPage{Page: page}

	if bh.head == nil {
		bh.head = newPage
	} else {
		bh.tail.next = newPage
		newPage.prev = bh.tail
	}

	bh.tail = newPage

	bh.Current = newPage
	bh.length++
}

func (bh *BrowserHistory) Pop() {
	if bh.length < 1 {
		return
	}

	if bh.length == 1 {
		bh.head = nil
		bh.tail = nil
		return
	}

	secLastPage := bh.tail.prev

	// remove last Page
	secLastPage.next = nil

	// secLastPage is now last Page
	bh.tail = secLastPage

	bh.Current = secLastPage
	bh.length--
}

func (bh *BrowserHistory) Forward() {
	cp := bh.Current
	if cp.next != nil {
		bh.Current = cp.next
	}
}

func (bh *BrowserHistory) Back() {
	cp := bh.Current
	if cp.prev != nil {
		bh.Current = cp.prev
	}
}

func (bh *BrowserHistory) Go(step int) {
	targetPage := bh.Current

	if step < 0 {
		// go backwards
		for i := 0; i > step && targetPage.prev != nil; i-- {
			targetPage = targetPage.prev
		}

		bh.Current = targetPage
		return
	}

	// go forwards
	for i := 0; i < step && targetPage.next != nil; i++ {
		targetPage = targetPage.next
	}

	bh.Current = targetPage
}
