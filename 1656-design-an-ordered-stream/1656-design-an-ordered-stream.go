type OrderedStream struct {
    pointer int
    list []string
}


func Constructor(n int) OrderedStream {
    return OrderedStream{
        pointer : 0,
        list : make([]string, n),
    }
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.list[idKey-1] = value
    chunk := make([]string,0)
    for {
        if this.pointer == len(this.list) {
            break
        }
        if this.list[this.pointer] != "" {
            chunk = append(chunk,this.list[this.pointer])
            this.pointer++
        } else {
            break
        }
    }
    return chunk
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */