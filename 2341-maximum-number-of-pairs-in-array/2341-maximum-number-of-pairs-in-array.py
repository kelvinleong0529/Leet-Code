class Solution:
    def numberOfPairs(self, nums: List[int]) -> List[int]:
        hash_map = {}
        for num in nums:
            if num in hash_map:
                hash_map[num] += 1
            else:
                hash_map[num] = 1
        
        pairs, leftover = 0, 0
        
        for key in hash_map:
            pairs += hash_map[key] // 2
            leftover += hash_map[key] % 2
        
        return [pairs,leftover]