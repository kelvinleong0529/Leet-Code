func maximizeSum(nums []int, k int) int {
    var answer, max int
    for i, v := range nums {
        if i == 0 {
            max = v
            continue
        }
        if v > max {
            max = v
        }
    }
    for i := 0; i < k; i ++ {
        answer += max
        max++
    }
    return answer
}