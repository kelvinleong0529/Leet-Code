func largestAltitude(gain []int) int {
    var ans = 0
    var current = 0
    for i := range gain {
        current = current + gain[i]
        if current > ans {
            ans = current
        }
    }
    return ans
}