func removeDuplicates(nums []int) int {
    hashMap := make(map[int]bool)

    for i := len(nums) - 1; i >= 0; i-- {
        _, ok := hashMap[nums[i]]
        if ok {
            nums = append(nums[:i], nums[i+1:]...)
        } else {
            hashMap[nums[i]] = true
        }
    }

    return len(hashMap)
}