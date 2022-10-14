class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        hash_map = {}
        ans = 0
        
        for num in nums:
            if num in hash_map:
                hash_map[num] += 1
            else:
                hash_map[num] = 1
        
        for key in hash_map:
            if key+k in hash_map:
                ans += hash_map[key]*hash_map[key+k]
        
        return ans