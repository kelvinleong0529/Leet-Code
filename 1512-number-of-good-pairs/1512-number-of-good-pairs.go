func numIdenticalPairs(nums []int) int {
    hashMap := make(map[int]int)
    answer := 0
    
    for i := range nums {
        _, ok := hashMap[nums[i]]
        if !ok {
            hashMap[nums[i]] = 0
        }
        hashMap[nums[i]]++
        if hashMap[nums[i]] > 1 {
            answer += hashMap[nums[i]] - 1
        }
    }
    
    return answer
}