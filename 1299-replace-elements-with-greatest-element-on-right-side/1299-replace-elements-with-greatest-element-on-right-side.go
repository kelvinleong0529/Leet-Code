func replaceElements(arr []int) []int {
    var ans = make([]int, len(arr))
    var max int = -1
    for i := len(arr) -1; i >=0 ; i-- {
        ans[i] = max
        if arr[i] > max {
            max = arr[i]
        }
    }
    return ans
}