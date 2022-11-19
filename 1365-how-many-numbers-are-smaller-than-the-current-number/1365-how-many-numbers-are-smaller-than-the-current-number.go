import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
    sortedNums := make([]int, len(nums))
    hashMap := make(map[int]int)
    copy(sortedNums, nums)
    sort.Ints(sortedNums)
    
    for i:= range sortedNums {
        _, ok := hashMap[sortedNums[i]]
        if !ok {
            hashMap[sortedNums[i]] = i
        }
    }
    
    for i := range nums {
        nums[i] = hashMap[nums[i]]
    }

    return nums
}