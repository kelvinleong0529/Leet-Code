class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        ans = []
        previous = None
        for i in range(len(nums)):
            if previous == None:
                previous = nums[i]
            else:
                if previous == nums[i]:
                    continue
                previous = nums[i]
            if previous > 0:
                break
            two_sum = self.twoSum(nums[i+1:],previous,-previous)
            if len(two_sum) > 0:
                for arr in two_sum:
                    ans.append(arr)
        return ans
                
            
    def twoSum(self, nums:List[int], element:int, target:int) -> List[int]:
        l, r = 0, len(nums) - 1
        ans = []
        while l < r:
            if nums[l] + nums[r] == target:
                if [element,nums[l],nums[r]] not in ans:
                    ans.append([element,nums[l],nums[r]])
                l += 1
                r -= 1
                continue
            if nums[l] + nums[r] > target:
                r -= 1
            else:
                l += 1
        return ans