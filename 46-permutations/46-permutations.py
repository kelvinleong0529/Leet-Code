class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        
        ans = []
        stack = []
        
        def backtracking(arr):
            if len(arr) == 0:
                ans.append(stack.copy())
                return
            
            for i in range(len(arr)):
                stack.append(arr[i])
                backtracking(arr[:i] + arr[i+1:])
                stack.pop()
            
        backtracking(nums)
        return ans