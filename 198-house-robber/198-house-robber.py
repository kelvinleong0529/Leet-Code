class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        if len(nums) == 2:
            return max(nums)
        ans = 0
        first, second = nums[0], max(nums[0],nums[1])
        for i in range(2,len(nums)):
            ans = max(nums[i]+first,second)
            first, second = second, ans
        return ans
        