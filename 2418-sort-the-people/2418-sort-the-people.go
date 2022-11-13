import "container/heap"

type Person struct {
    name string
    height int
}

type maxHeap []Person

func (m maxHeap) Len() int { return len(m) }
func (m maxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m maxHeap) Less(i, j int) bool { return m[i].height > m[j].height }

func (m *maxHeap) Push(x interface{}) {
    *m = append(*m, x.(Person))
} 

func (m *maxHeap) Pop() interface{} {
    old := *m
    n := len(old)
    last := old[n-1]
    *m = old[0:n-1]
    return last
}

func sortPeople(names []string, heights []int) []string {
    h := &maxHeap{}
    
    for i := 0; i < len(names); i++ {
        p := Person{
            name : names[i],
            height : heights[i],
        }
        *h = append(*h, p)
    }
    
    heap.Init(h)
    var ans = make([]string, len(names))
    index := 0
    for h.Len() > 0 {
        pop := heap.Pop(h).(Person)
        ans[index] = pop.name
        index++
    }
    return ans
}