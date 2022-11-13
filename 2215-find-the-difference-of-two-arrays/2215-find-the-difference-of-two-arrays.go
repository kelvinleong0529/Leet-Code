type stats struct {
    first bool
    second bool
}

func (s *stats) update(i int) {
    if i == 1 {
        s.first = true
    } else {
        s.second = true
    }
}

func findDifference(nums1 []int, nums2 []int) [][]int {
    
    hashMap := make(map[int]*stats)
    for i := range nums1 {
        if _, ok := hashMap[nums1[i]]; !ok {
            hashMap[nums1[i]] = &stats{}
        }
        hashMap[nums1[i]].update(1)
    }
    for i := range nums2 {
        if _, ok := hashMap[nums2[i]]; !ok {
            hashMap[nums2[i]] = &stats{}
        }
        hashMap[nums2[i]].update(2)
    }
    var ans = make([][]int,2)
    ans[0] = make([]int,0)
    ans[1] = make([]int,0)
    for k, v := range hashMap {
        if !(v.first && v.second) {
            if !v.first {
                ans[1] = append(ans[1], k)
            }
            if !v.second {
                ans[0] = append(ans[0], k)
            }
        }
    }
    return ans
}