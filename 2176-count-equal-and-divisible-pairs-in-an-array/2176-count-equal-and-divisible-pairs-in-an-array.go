func countPairs(nums []int, k int) int {
    hashMap := make(map[int][]int)
    ans := 0
    for i, v := range nums {
        if _, ok := hashMap[v]; ok {
            hashMap[v] = append(hashMap[v], i)
        } else {
            hashMap[v] = []int{i}
        }
        if i % k == 0 {
            ans = ans + len(hashMap[v]) - 1
        } else {
            for _, value := range hashMap[v] {
                if value == i {
                    continue
                }
                if value*i % k == 0 {
                    ans++
                }
            }
        }
    }
    return ans
}