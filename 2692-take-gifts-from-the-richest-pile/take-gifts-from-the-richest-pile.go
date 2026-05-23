type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
    h := &MaxHeap{}
	heap.Init(h)

	for _, gift := range gifts {
		heap.Push(h, gift)
	}

	for k > 0 {
		current := float64(heap.Pop(h).(int))
		sqr := math.Floor(math.Sqrt(current))
		heap.Push(h, int(sqr))
		k--
	}

	var ans int64

	for h.Len() > 0 {
		current :=  float64(heap.Pop(h).(int))
		ans += int64(current)
	}

	return ans
}