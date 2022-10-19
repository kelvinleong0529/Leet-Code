func divideArray(nums []int) bool {
    hashMap := make(map[int]int)
    
    for i := range nums {
        _, ok := hashMap[nums[i]]
        if ok {
            hashMap[nums[i]]++
        } else {
            hashMap[nums[i]] = 1
        }
    }
    
    for _, v := range hashMap {
        if v%2 == 1 {
            return false
        }
    }
    return true
}