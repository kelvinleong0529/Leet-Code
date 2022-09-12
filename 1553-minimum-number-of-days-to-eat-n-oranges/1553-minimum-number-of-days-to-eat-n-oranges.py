class Solution:
    def minDays(self, n: int) -> int:
        da = {0:0,1:1}

        def dfs(i):
            if i in da:
                return da[i]
            
            result1 = 1 + (i%2) + dfs(i//2)
            result2 = 1 + (i%3) + dfs(i//3)

            if result1 <= result2:
                da[i] = result1
            else:
                da[i] = result2
            
            return da[i]

        return dfs(n)