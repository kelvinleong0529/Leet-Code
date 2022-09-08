class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        pointer = len(nums) - 2
        count = 1
        for i in range(2,len(nums)+1):
            if nums[len(nums)-i] >= count:
                pointer -= count
                count = 1
                continue
            count += 1

        if pointer < 0:
            return True
        else:
            return False
                