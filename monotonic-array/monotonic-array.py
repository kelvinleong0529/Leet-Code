class Solution:
    def isMonotonic(self, nums: List[int]) -> bool:
        copy = nums.copy()
        copy.sort()
        return (nums == copy or nums == copy[::-1])
            