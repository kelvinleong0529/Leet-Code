func leftRightDifference(nums []int) []int {
    leftSum := make([]int, len(nums))
    sum := 0
    
    for i := range nums {
        if i == len(nums)-1 {
            sum += nums[i]
            break
        }
        leftSum[i+1] = sum+nums[i]
        sum = leftSum[i+1]
    }
    
    result := make([]int, len(nums))
    for i := range leftSum {
        rightSum := sum - leftSum[i] - nums[i]
        diff := rightSum - leftSum[i]
        if diff < 0 {
            diff = -diff
        }
        result[i] = diff
    }
    
    return result
}