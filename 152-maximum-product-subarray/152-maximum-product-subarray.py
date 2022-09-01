class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        ans, cur_min, cur_max = 0, 0, 0
        for i in nums:
            if i == 0:
                cur_min, cur_max = 1, 1
                continue
            if i < 0:
                temp_min, temp_max = cur_max, cur_min
            else:
                temp_min, temp_max = cur_min, cur_max
            cur_min, cur_max = min(i * temp_min, i), max(i*temp_max, i)
            if cur_max > ans:
                ans = cur_max
        return ans