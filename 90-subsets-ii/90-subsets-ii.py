class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        ans = []
        stack = []
        
        def backtrack(i):
            if i >= len(nums):
                temp = stack.copy()
                temp.sort()
                if temp not in ans:
                    ans.append(temp)
                return
                    
            stack.append(nums[i])
            backtrack(i+1)
            stack.pop()
            
            backtrack(i+1)
            
        backtrack(0)
        return ans