class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        ans = []
        
        def backtrack(start,stack):
            if len(stack) == k:
                ans.append(stack.copy())
                return
            
            for i in range(start,n+1):
                stack.append(i)
                backtrack(i+1,stack)
                stack.pop()
                
        backtrack(1,[])
        return ans