func findDisappearedNumbers(nums []int) []int {
    hashMap := make(map[int]bool)
    for i := 1; i < len(nums)+1; i++ {
        hashMap[i] = false
    }
    for i := range nums {
        _, ok := hashMap[nums[i]]
        if ok {
            hashMap[nums[i]] = true
        }
    }
    ans := make([]int,0)
    for k, v := range hashMap {
        if !v {
            ans = append(ans, k)
        }
    }
    return ans
}