package main

import "fmt"

func reverseInteger(x int) int {
	var ans int
	for {
		if x <= 0 {
			break
		}
		remainder := x % 10
		ans = ans*10 + remainder
		x = x / 10
	}
	return ans
}

func countDistinctIntegers(nums []int) int {
	hashMap := make(map[int]bool)
	for i := range nums {
		_, ok := hashMap[nums[i]]
		if !ok {
			hashMap[nums[i]] = true
		}
		reversed := reverseInteger(nums[i])
		_, ok = hashMap[reversed]
		if !ok {
			hashMap[reversed] = true
		}
	}
	return len(hashMap)
}
