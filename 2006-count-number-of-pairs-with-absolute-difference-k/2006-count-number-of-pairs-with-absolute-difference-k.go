func countKDifference(nums []int, k int) int {
    hashMap := make(map[int]int)
    ans := 0
    for i := range nums {
        add := nums[i]+k
        minus := nums[i]-k
        if v, ok := hashMap[add]; ok {
            ans += v
        }
        if v, ok := hashMap[minus]; ok {
            ans += v
        }
        if v, ok := hashMap[nums[i]]; ok {
            v++
            hashMap[nums[i]] = v
        } else {
            hashMap[nums[i]] = 1
        }
    }
    return ans
}