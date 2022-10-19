func numberOfPairs(nums []int) []int {
    answer := make([]int,2)
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
        answer[0] += v/2
        answer[1] += v%2
    }
    return answer
}