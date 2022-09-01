class Solution:
    def climbStairs(self, n: int) -> int:
        da = [1,1]
        count = len(da)
        while count <= n:
            da.append(da[count-2]+da[count-1])
            count += 1
        return da[n]
            