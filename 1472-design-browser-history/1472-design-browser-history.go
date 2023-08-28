
type BrowserListNode struct {
	val  string
	prev *BrowserListNode
	next *BrowserListNode
}

type BrowserHistory struct {
	cur *BrowserListNode
}

func newBrowserListNode(val string) *BrowserListNode {
	return &BrowserListNode{val: val}
}

func Constructor(homepage string) BrowserHistory {
	BrowserListNode := newBrowserListNode(homepage)
	return BrowserHistory{
		cur: BrowserListNode,
	}
}

func (this *BrowserHistory) Visit(url string) {
	BrowserListNode := newBrowserListNode(url)
	BrowserListNode.prev = this.cur
	this.cur.next = BrowserListNode
	this.cur = this.cur.next
}

func (this *BrowserHistory) Back(steps int) string {
	for {
		if this.cur.prev == nil || steps <= 0 {
			break
		}
		this.cur = this.cur.prev
		steps--
	}
	return this.cur.val
}

func (this *BrowserHistory) Forward(steps int) string {
	for {
		if this.cur.next == nil || steps <= 0 {
			break
		}
		this.cur = this.cur.next
		steps--
	}
	return this.cur.val
}
