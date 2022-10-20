func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    answer := make([]int,2)
    for i,v := range nums {
        diff := target - v
        _, ok := hashMap[diff]
        if ok {
            answer = []int{hashMap[diff],i}
            break
        } else {
            hashMap[v] = i
        }
    }
    return answer
}