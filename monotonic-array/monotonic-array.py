class Solution:
    def isMonotonic(self, nums: List[int]) -> bool:
        if len(nums)==1:
            return True
        temp1 = nums[1:]
        temp2 = nums[:-1]
        ans = []
        for i in range (len(temp1)):
            ans.append(temp1[i]-temp2[i])
        if max(ans)>0 and min(ans)<0:
            return False
        else:
            return True