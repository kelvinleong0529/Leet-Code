func fourSum(nums []int, target int) [][]int {
    slices.Sort(nums)
    ans := [][]int{}
    // Use length 0, capacity 4
    quad := make([]int, 0, 4)

    var recursive func(k int, start int, target int)
    recursive = func(k int, start int, target int) {
        if k > 2 {
            for i := start; i <= len(nums)-k; i++ {
                // Skip duplicates
                if i > start && nums[i] == nums[i-1] {
                    continue
                }
                quad = append(quad, nums[i])
                // Pass i+1 to move to the next unique element
                recursive(k-1, i+1, target-nums[i])
                quad = quad[:len(quad)-1]
            }
            return
        }

        // Standard 2Sum Two-Pointer logic
        left, right := start, len(nums)-1
        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                // Create a clean copy of the quad + two pointers
                row := make([]int, 0, 4)
                row = append(row, quad...)
                row = append(row, nums[left], nums[right])
                ans = append(ans, row)
                
                left++
                // Skip duplicates for 'left'
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }

    recursive(4, 0, target)
    return ans
}