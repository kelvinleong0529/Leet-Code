class Solution:
    def divideArray(self, nums: List[int]) -> bool:
        hash_map = {}
        
        for num in nums:
            if num in hash_map:
                hash_map[num] += 1
            else:
                hash_map[num] = 1
        
        for key in hash_map:
            if hash_map[key] % 2:
                return False
        
        return True