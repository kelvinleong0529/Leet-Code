class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if len(nums) == 0:
            return []
        start = current = nums[0]
        ans = []
        for i in range(1,len(nums)):
            v = nums[i]
            if current + 1 != v:
                if start == current:
                    ans.append(str(start))
                else:
                    ans.append(str(start) + "->" + str(current))
                start = current = v
            if v > current:
                current = v
        if start == current:
            ans.append(str(start))
        else:
            ans.append(str(start) + "->" + str(current))
        return ans