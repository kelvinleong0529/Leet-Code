func arithmeticTriplets(nums []int, diff int) int {
    var answer int
    nums1 := make([]int, len(nums))
    nums2 := make([]int, len(nums))
    hashMap := make(map[int]bool)
    for i, v := range nums {
        hashMap[v] = true
        nums1[i] = v - diff
        nums2[i] = v - 2 * diff
    }
    for i := range nums {
        _, ok1 := hashMap[nums[i]]
        _, ok2 := hashMap[nums1[i]]
        _, ok3 := hashMap[nums2[i]]
        if ok1 && ok2 && ok3 {
            answer++
        }
    }
    return answer
}