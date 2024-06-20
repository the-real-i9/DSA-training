package webBrowser

type WebPage struct {
	Title string
	Body  string
}

type browserPage struct {
	prev *browserPage
	Main WebPage
	next *browserPage
}

type History struct {
	head    *browserPage
	tail    *browserPage
	Current *browserPage
	length  int
}

func (bh *History) Push(page WebPage) {
	newBrowserPage := &browserPage{Main: page}

	if bh.head == nil {
		bh.head = newBrowserPage
	} else {
		bh.tail.next = newBrowserPage
		newBrowserPage.prev = bh.tail
	}

	bh.tail = newBrowserPage

	bh.Current = newBrowserPage
	bh.length++
}

func (bh *History) Pop() {
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

func (bh *History) Forward() {
	cp := bh.Current
	if cp.next != nil {
		bh.Current = cp.next
	}
}

func (bh *History) Back() {
	cp := bh.Current
	if cp.prev != nil {
		bh.Current = cp.prev
	}
}

func (bh *History) Go(step int) {
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
