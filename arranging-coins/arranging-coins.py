import math
class Solution:
    def arrangeCoins(self, n: int) -> int:
        if n ==0:
            return 0
        ans = math.floor((n*2)**0.5)
        if ans*(ans+1)*0.5>n:
            return ans-1
        else:
            return ans
