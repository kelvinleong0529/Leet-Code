class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        ans = [1]
        cur = 1
        while cur < 2**31 - 1:
            ans.append(int(cur*4))
            ans.append(float(1/(cur*4)))
            cur = cur*4
        return n in ans