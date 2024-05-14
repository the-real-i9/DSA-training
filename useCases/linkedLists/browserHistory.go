package lluc

type webPage struct {
	Title string
	Body  string
}

type browserPage struct {
	prev *browserPage
	Page webPage
	next *browserPage
}

type browserHistory struct {
	head    *browserPage
	tail    *browserPage
	Current *browserPage
	length  int
}

func (bh *browserHistory) Push(page webPage) {
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

func (bh *browserHistory) Pop() {
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

func (bh *browserHistory) Forward() {
	cp := bh.Current
	if cp.next != nil {
		bh.Current = cp.next
	}
}

func (bh *browserHistory) Back() {
	cp := bh.Current
	if cp.prev != nil {
		bh.Current = cp.prev
	}
}

func (bh *browserHistory) Go(step int) {
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
