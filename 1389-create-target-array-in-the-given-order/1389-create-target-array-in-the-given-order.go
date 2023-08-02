func createTargetArray(nums []int, index []int) []int {
    target := make([]int, 0, len(nums))
    for i := range nums {
        targetNum := nums[i]
        targetIndex := index[i]
        if targetIndex > len(target) {
            target = append(target, targetNum)
        } else {
            target = append(target[:targetIndex+1], target[targetIndex:]...)
            target[targetIndex] = targetNum
        }
    } 
    return target
}