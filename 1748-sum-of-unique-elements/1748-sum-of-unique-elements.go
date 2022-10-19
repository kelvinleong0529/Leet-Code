func sumOfUnique(nums []int) int {
    hashMap := make(map[int]int)
    
    for i := range nums {
        _, ok := hashMap[nums[i]]
        if ok {
            hashMap[nums[i]]++
        } else {
            hashMap[nums[i]] = 1
        }
    }
    
    answer := 0
    for k, v := range hashMap {
        if v == 1 {
            answer += k
        }
    }
    
    return answer
}