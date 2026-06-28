func maxNumber(nums1 []int, nums2 []int, k int) []int {
    m, n := len(nums1), len(nums2)
    var best []int

    for i := max(0, k-n); i <= min(k, m); i++ {
        sub1 := maxSubsequence(nums1, i)
        sub2 := maxSubsequence(nums2, k-i)
        merged := merge(sub1, sub2)
        if compare(merged, 0, best, 0) > 0 {
            best = merged
        }
    }
    return best
}

// pick `k` digits from nums in order, maximizing the number
func maxSubsequence(nums []int, k int) []int {
    drop := len(nums) - k  // how many we're allowed to drop
    stack := []int{}
    for _, d := range nums {
        // pop if current digit is larger and we still have drops left
        for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
            stack = stack[:len(stack)-1]
            drop--
        }
        stack = append(stack, d)
    }
    return stack[:k]
}

// merge two subsequences greedily, always picking the larger head
func merge(s1, s2 []int) []int {
    res := make([]int, 0, len(s1)+len(s2))
    i, j := 0, 0
    for i < len(s1) || j < len(s2) {
        if compare(s1, i, s2, j) >= 0 {
            res = append(res, s1[i])
            i++
        } else {
            res = append(res, s2[j])
            j++
        }
    }
    return res
}

// lexicographic compare starting at given offsets
func compare(a []int, i int, b []int, j int) int {
    for i < len(a) && j < len(b) {
        if a[i] != b[j] {
            return a[i] - b[j]
        }
        i++
        j++
    }
    return (len(a) - i) - (len(b) - j)  // longer one wins if prefix equal
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}