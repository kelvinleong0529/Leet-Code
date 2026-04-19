func permute(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int {{}}
    }

    perms := permute(nums[1:])
    ans := make([][]int, 0)
    for _, v := range perms {
        for i := range len(v) + 1 {
            deepCopy := make([]int, len(v))
            copy(deepCopy, v)
            permutation := slices.Insert(deepCopy, i, nums[0])
            ans = append(ans, permutation)
        }
    }

    return ans
}