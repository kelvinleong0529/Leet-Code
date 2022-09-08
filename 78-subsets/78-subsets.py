class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = []
        stack = []
        
        def backtrack(i):
            if i >= len(nums):
                ans.append(stack.copy())
                return
            
            # to include
            stack.append(nums[i])
            backtrack(i+1)
            stack.pop()
            
            # to NOT include
            backtrack(i+1)
            
        backtrack(0)
        return ans