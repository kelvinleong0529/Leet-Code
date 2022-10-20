func decompressRLElist(nums []int) []int {
    answer := make([]int,0)
    
    for i := 0; i < len(nums); i += 2 {
        for j := 0; j < nums[i]; j++ {
            answer = append(answer,nums[i+1])
        }
    }
    
    return answer
}