func minOperations(nums []int) int {
    answer := 0
    current := 0
    
    for i := range nums {
        if current == 0 {
            current = nums[i]
        } else {
            if current > nums[i] {
                answer += current - nums[i]
            } else {
                current = nums[i]
            }
        }
        current++
    }
    return answer
}