type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type LRUCache struct {
	capacity int
	hashMap  map[int]*Node
	left     *Node
	right    *Node
}

func Constructor(capacity int) LRUCache {
	left, right := &Node{}, &Node{}
	left.next, right.prev = right, left
	return LRUCache{
		capacity: capacity,
		hashMap:  make(map[int]*Node, 0),
		left:     left,
		right:    right,
	}
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) insert(node *Node) {
	right := this.right.prev
	right.next = node
	node.prev, node.next = right, this.right
	this.right.prev = node
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	this.remove(v)
	this.insert(v)
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hashMap[key]; ok {
		this.remove(v)
	}

	node := &Node{key: key, value: value}
	this.insert(node)
	this.hashMap[key] = node

	if len(this.hashMap) > this.capacity {
		delete(this.hashMap, this.left.next.key)
		this.remove(this.left.next)
	}
}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */