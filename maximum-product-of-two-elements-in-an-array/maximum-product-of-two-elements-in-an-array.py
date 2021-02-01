class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        temp = nums.pop(nums.index(max(nums)))
        return (temp-1)*(max(nums)-1)