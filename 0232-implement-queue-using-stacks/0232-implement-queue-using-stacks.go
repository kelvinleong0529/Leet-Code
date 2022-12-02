type MyQueue struct {
    length int
    array []int
}


func Constructor() MyQueue {
    return MyQueue{
        length : 0,
        array : make([]int,0),
    }
}


func (this *MyQueue) Push(x int)  {
    this.length++
    this.array = append(this.array, x)
}


func (this *MyQueue) Pop() int {
    first := this.array[0]
    this.array = this.array[1:]
    this.length--
    return first
}


func (this *MyQueue) Peek() int {
    return this.array[0]
}


func (this *MyQueue) Empty() bool {
    return this.length == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */