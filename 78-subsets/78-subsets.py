class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:

        ans = []
        stack = []
        
        def backtracking(i):
            if i == len(nums):
                ans.append(stack.copy())
                return
            
            # NOT to include
            backtracking(i+1)
            
            # to include
            stack.append(nums[i])
            backtracking(i+1)
            stack.pop()
            
        backtracking(0)
        
        return ans