func unequalTriplets(nums []int) int {
    hashMap := make(map[int]int)
    for _, v := range nums {
        hashMap[v]++
    }
    array := make([]int, len(hashMap))
    index := 0
    for _, v := range hashMap {
        array[index] = v
        index++
    }
    left := 0
    right := len(nums)
    var answer int
    for _, v := range array {
        right -= v
        answer += left * v * right
        left += v
    }
    return answer
}