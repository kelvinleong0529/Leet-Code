class Solution(object):
    def canPartition(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        if len(nums) == 1:
            return False
        if sum(nums) % 2 != 0:
            return False
        target = sum(nums)//2
        my_set = set()
        for i in nums:
            if i == target:
                return True
            for j in my_set.copy():
                temp = j + i
                if temp == target:
                    return True
                my_set.add(temp)
            my_set.add(i)
        return False