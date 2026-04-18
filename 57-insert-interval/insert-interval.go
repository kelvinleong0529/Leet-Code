func insert(intervals [][]int, newInterval []int) [][]int {
    ans := make([][]int, 0)
    
    for i, v := range intervals {
        if newInterval[1] < v[0] {
            ans = append(ans, newInterval)
            ans = append(ans, intervals[i:]...)
            return ans
        } else if newInterval[0] > v[1] {
            ans = append(ans, v)
        } else {
            newInterval[0] = min(newInterval[0], v[0])
            newInterval[1] = max(newInterval[1], v[1])
        }
    }

    ans = append(ans, newInterval)

    return ans
}