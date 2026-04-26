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

func nthUglyNumber(n int) int {
	h := &MinHeap{1}
	heap.Init(h)
	prev := 0
	count := 1

	factors := []int{2, 3, 5}

	for {
		top := heap.Pop(h).(int)
		if top == prev {
			continue
		}
		if count == n {
			return top
		}
		count++
		prev = top
		for i := range factors {
			heap.Push(h, top * factors[i])
		}
	}
}