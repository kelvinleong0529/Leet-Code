class Solution:
    def minStartValue(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            if not i:
                continue
            nums[i] = nums[i-1] + nums[i]
        return max(1-min(nums),1)