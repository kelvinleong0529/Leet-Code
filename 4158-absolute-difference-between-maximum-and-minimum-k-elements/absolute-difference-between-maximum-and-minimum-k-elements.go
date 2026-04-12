func absDifference(nums []int, k int) int {
    slices.Sort(nums)

    var largest int
    var smallest int

    for i := len(nums) - 1; i >= len(nums) - k; i-- {
        largest += nums[i]
    }

    for i := 0; i <= k-1; i++ {
        smallest += nums[i]
    }

    return largest - smallest
}