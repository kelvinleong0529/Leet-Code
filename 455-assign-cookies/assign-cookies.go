import (
	"container/heap"
)

// MinHeap is a type that implements heap.Interface
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] } // min-heap
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[0 : n-1]
	return val
}

func findContentChildren(g []int, s []int) int {
	greedHeap := new(MinHeap)
	heap.Init(greedHeap)
	for i := range g {
		heap.Push(greedHeap, g[i])
	}

	cookiesHeap := new(MinHeap)
	heap.Init(cookiesHeap)
	for i := range s {
		heap.Push(cookiesHeap, s[i])
	}

	result := 0
	for {
		if cookiesHeap.Len() == 0 || greedHeap.Len() == 0 {
			break
		}
		cookie := heap.Pop(cookiesHeap).(int)
		greed := (*greedHeap)[0]
		if cookie >= greed {
			heap.Pop(greedHeap)
            result++
		}
	}

	return result
}