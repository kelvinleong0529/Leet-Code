class Solution:
    def distinctAverages(self, nums: List[int]) -> int:
        nums.sort()
        hashMap = {}
        while len(nums) > 0:
            first, last = nums[0], nums[-1]
            nums = nums[1:-1]
            avg = (first+last)/2
            if avg not in hashMap:
                hashMap[avg] = True
        
        return len(hashMap)
        