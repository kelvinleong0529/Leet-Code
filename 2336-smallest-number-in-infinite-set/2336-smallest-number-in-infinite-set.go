type SmallestInfiniteSet struct {
	smallestNumber int
	newlyAdded     []int
	numbersAdded   map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		smallestNumber: 1,
		newlyAdded:     make([]int, 0),
		numbersAdded:   make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var smallest int
	if len(this.newlyAdded) > 0 {
		smallest = this.newlyAdded[0]
		this.newlyAdded = this.newlyAdded[1:]
		delete(this.numbersAdded, smallest)
	} else {
		smallest = this.smallestNumber
		this.smallestNumber++
	}
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.numbersAdded[num] {
		return
	}
	if num < this.smallestNumber {
		this.newlyAdded = append(this.newlyAdded, num)
		this.numbersAdded[num] = true
		sort.Slice(this.newlyAdded, func(i, j int) bool {
			return this.newlyAdded[i] < this.newlyAdded[j]
		})
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */