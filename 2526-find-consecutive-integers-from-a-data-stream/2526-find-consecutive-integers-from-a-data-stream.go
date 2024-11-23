type DataStream struct {
	value int
	streak int
	currentStreak int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value: value,
		streak: k,
	}
}

func (this *DataStream) Consec(num int) bool {
	if num != this.value {
		this.currentStreak = 0
		return false
	}
	this.currentStreak++
	return this.currentStreak >= this.streak
}


/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */