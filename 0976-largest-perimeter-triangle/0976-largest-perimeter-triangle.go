func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    reverseIntSlice(nums)
        
    for i := 0; i < len(nums)-2; i++ {
        if nums[i] < nums[i+1] + nums[i+2] {
            return nums[i] + nums[i+1] + nums[i+2]
        }
    }
    return 0
}

func reverseIntSlice(arr []int) {
    n := len(arr)

    for i := 0; i < n/2; i++ {
        arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
    }
}