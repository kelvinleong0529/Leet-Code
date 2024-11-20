type FrequencyTracker struct {
	numberMap    map[int]int
	frequencyMap map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		numberMap:    make(map[int]int),
		frequencyMap: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	this.numberMap[number]++
	frequency := this.numberMap[number]
	this.frequencyMap[frequency]++
	if frequency != 1 {
		this.frequencyMap[frequency-1]--
	}
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if _, ok := this.numberMap[number]; ok {
		this.numberMap[number]--
		frequency := this.numberMap[number]
		this.frequencyMap[frequency+1]--
		if frequency == 0 {
			delete(this.numberMap, number)
		} else {
			this.frequencyMap[frequency]++
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	val, ok := this.frequencyMap[frequency]
	return ok && val > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */