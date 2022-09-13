class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        stack = [] # [val, index]
        ans = [-1] * len(nums)
        
        for i,n in enumerate(nums*2):
            while stack and n > stack[-1][0]:
                _ , stackI = stack.pop()
                ans[stackI] = n
            if i < len(nums):
                stack.append([n,i])
        
        return ans