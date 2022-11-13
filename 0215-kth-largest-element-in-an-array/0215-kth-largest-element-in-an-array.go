import "container/heap"

type maxHeap []int

func (m maxHeap) Len() int { return len(m) }
func (m maxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m maxHeap) Less(i, j int) bool {return m[i] > m[j] }

func (m *maxHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
    old := *m
    n := len(old)
    last := old[n-1]
    *m = old[0:n-1]
    return last
}

func findKthLargest(nums []int, k int) int {
    h := &maxHeap{}
    heap.Init(h)
    for i := range nums {
        heap.Push(h, nums[i])
    }
    k--
    for k > 0 {
        heap.Pop(h)
        k--
    }
    return heap.Pop(h).(int)
}