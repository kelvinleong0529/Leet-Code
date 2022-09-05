class Solution:
    def minOperations(self, nums: List[int], x: int) -> int:
        target = sum(nums) - x
        if target == 0:
            return len(nums)
        hash_map = {0:0}
        ans = total = 0
        for i,v in enumerate(nums):
            total += v
            diff = total - target
            if diff in hash_map:
                temp = i+1-hash_map[diff]
                if temp > ans:
                    ans = temp
            hash_map[total] = i+1
        if ans == 0:
            return -1
        else:
            return len(nums)-ans
                