type MyHashSet struct {
    hashMap map[int]bool
}


func Constructor() MyHashSet {
    return MyHashSet{
        hashMap : make(map[int]bool),
    }
}


func (this *MyHashSet) Add(key int)  {
    this.hashMap[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    if _, ok := this.hashMap[key]; ok {
        delete(this.hashMap, key)
    }
}


func (this *MyHashSet) Contains(key int) bool {
    if _, ok := this.hashMap[key]; ok {
        return true
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */