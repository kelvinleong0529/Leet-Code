class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        hash_map = {}
        ans = 0
        for num in nums:
            if num in hash_map:
                hash_map[num] += 1
                ans += hash_map[num]-1
            else:
                hash_map[num] = 1
            
        return ans