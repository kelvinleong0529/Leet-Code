type Pair struct {
	sum    int
	index1 int
	index2 int
}

type PairMinHeap []Pair

func (m PairMinHeap) Len() int {
	return len(m)
}

func (m PairMinHeap) Less(i, j int) bool {
	return m[i].sum < m[j].sum
}

func (m PairMinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *PairMinHeap) Push(x any) {
	*m = append(*m, x.(Pair))
}

func (m *PairMinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func initHeap(pairs []Pair) *PairMinHeap {
	h := &PairMinHeap{}
	heap.Init(h)

	for _, pair := range pairs {
		heap.Push(h, pair)
	}

	return h
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	arr := make([]Pair, len(nums1))

	for i := range nums1 {
		arr[i] = Pair{sum: nums1[i] + nums2[0], index1: i, index2: 0}
	}

	h := initHeap(arr)

	ans := make([][]int, k)

	for i := 0; i < k; i++ {
		pair := heap.Pop(h).(Pair)
		ans[i] = []int{nums1[pair.index1], nums2[pair.index2]}
		if pair.index2 < len(nums2) - 1 {
			heap.Push(h, 
            Pair{sum: 
                nums1[pair.index1] + nums2[pair.index2+1], 
                index1: pair.index1, 
                index2: pair.index2+1,
                })
		}
	}

	return ans
}