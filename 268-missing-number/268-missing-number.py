class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        length = len(nums)
        original_sum = int((length+1)*length/2)
        current_sum = 0
        for i in nums:
            current_sum += i
        return original_sum - current_sum