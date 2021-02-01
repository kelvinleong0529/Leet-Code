class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        ans = []
        for i in nums:
            if target-i in nums:
                if nums.index(i) != nums.index(target-i):
                    ans.append(nums.index(i))
                    ans.append(nums.index(target-i))
                    break
                else:
                    if nums.count(i) > 1 :
                        ans.append(nums.index(i))
                        nums.remove(i)
                        ans.append(nums.index(i)+1)
                        break
        return ans