class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        front1, back1 = nums[0], max(nums[0], nums[1])
        front2, back2 = nums[-1], max(nums[-1], nums[-2])
        max1, max2 = back1, back2
        for i in range(2, len(nums)-1):
            max1 = max(nums[i]+front1, back1)
            front1, back1 = back1, max1

            max2 = max(nums[len(nums)-i-1]+front2, back2)
            front2, back2 = back2, max2
            
        return max(max1,max2)