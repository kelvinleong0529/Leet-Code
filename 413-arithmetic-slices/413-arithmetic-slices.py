class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        if len(nums) < 3:
            return 0
        difference = []
        for i in range(len(nums)-1):
            difference.append(nums[i+1]-nums[i])
        ans = 0
        key = 1
        value = 0
        previous = None
        for i in difference:
            if i == previous:
                value += key
                key += 1
            else:
                previous = i
                ans += value
                value = 0
                key = 1
        return ans + value
            