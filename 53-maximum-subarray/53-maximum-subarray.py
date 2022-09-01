class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        current_max = -100001
        ans = -100001
        for i in range(len(nums)):
            current_max = max(current_max+nums[i], nums[i])
            if current_max > ans:
                ans = current_max
        return ans