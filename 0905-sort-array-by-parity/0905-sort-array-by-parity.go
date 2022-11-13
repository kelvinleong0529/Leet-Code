func sortArrayByParity(nums []int) []int {
    odd := make([]int, 0)
    even := make([]int, 0)
    ans := make([]int, 0)
    
    for i := range nums {
        if nums[i]%2 == 0 {
            even = append(even, nums[i])
        } else {
            odd = append(odd, nums[i])
        }
    }
    
    ans = append(ans , even...)
    ans = append(ans , odd...)
    return ans
}