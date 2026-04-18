func topKFrequent(nums []int, k int) []int {
    hashMap := make(map[int]int)

    for _, num := range nums {
        hashMap[num]++
    }

    count := make([][]int, len(nums))

    for key, v := range hashMap {
        if count[v-1] == nil {
            count[v-1] = make([]int, 0)
        }
        count[v-1] = append(count[v-1], key)
    }

    ans := make([]int, 0)

    for i := len(count) - 1; i >= 0 && k > 0; i-- {
        if count[i] != nil {
            ans = append(ans, count[i]...)
        }
        k = k - len(count[i])
    }

    return ans
}