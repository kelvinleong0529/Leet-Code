func findLonely(nums []int) []int {
    hashMap := make(map[int]int)
    
    for i := range nums {
        _, ok := hashMap[nums[i]]
        if ok {
            hashMap[nums[i]]++
        } else {
            hashMap[nums[i]] = 1
        }
    }
    
    ans := make([]int,0)
    for k, v := range hashMap {
        if v != 1 {
            continue
        }
        _, ok1 := hashMap[k-1]
        _, ok2 := hashMap[k+1]
        
        if !ok1 && !ok2 {
            ans = append(ans,k)
        }
    }
    
    return ans
}