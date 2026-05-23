type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func (h MinHeap) Peek() (int, bool) {
	if len(h) == 0 {
		return 0, false // Heap is empty
	}
	return h[0], true // The smallest element is always at index 0
}

type KthLargest struct {
	size    int
	maxHeap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return KthLargest{
		size:    k,
		maxHeap: h,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.maxHeap, val)
    if this.maxHeap.Len() > this.size {
        heap.Pop(this.maxHeap)
    }
	res, _ := this.maxHeap.Peek()
	return res
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */