import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Worker struct {
	ratio   float64
	quality int
}

// MaxHeap for qualities (we want to pop the largest quality when heap size > k)
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	workers := make([]Worker, n)

	// Build array of workers with their wage-to-quality ratio
	for i := 0; i < n; i++ {
		workers[i] = Worker{
			ratio:   float64(wage[i]) / float64(quality[i]),
			quality: quality[i],
		}
	}

	// Sort workers by ratio ascending
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})

	h := &MaxHeap{}
	heap.Init(h)

	totalQuality := 0
	minCost := math.Inf(1)

	for _, w := range workers {
		heap.Push(h, w.quality)
		totalQuality += w.quality

		if h.Len() > k {
			totalQuality -= heap.Pop(h).(int)
		}
		if h.Len() == k {
			cost := float64(totalQuality) * w.ratio
			if cost < minCost {
				minCost = cost
			}
		}
	}

	return minCost
}
