func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    hashMap := make(map[int][]bool)
    
    iterate := func (nums []int, index int) {
        for i := range nums {
            _, ok := hashMap[nums[i]]
            if !ok {
                hashMap[nums[i]] = []bool {false, false, false}
            }
            hashMap[nums[i]][index] = true
        }
    }
    
    count := func (arr []bool) int {
        count := 0
        for i := range arr {
            if arr[i] {
                count++
            }
        }
        return count
    }
    
    iterate(nums1,0)
    iterate(nums2,1)
    iterate(nums3,2)

    answer := make([]int,0)
    
    for k, v := range hashMap {
        if count(v) >= 2 {
            answer = append(answer,k)
        }
    }
    
    return answer
}