func arraySign(nums []int) int {
    answer := 1
    for i := range nums {
        if nums[i] == 0 {
            return 0
        }
        if nums[i] < 0 {
            answer = answer*-1
        }
    }
    return answer
}