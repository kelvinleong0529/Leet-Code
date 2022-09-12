class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        
        nums_set = {i for i in nums}
        ans = 0
        
        for n in nums:
            if (n-1) in nums_set:
                continue
            length = 1
            while (n+length) in nums_set:
                length += 1
            if length > ans:
                ans = length
                
        return ans