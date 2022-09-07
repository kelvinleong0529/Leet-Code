class Solution:
    def trailingZeroes(self, n: int) -> int:
        ans = 0
        exponent = 1
        while n // (5**exponent) > 0:
            ans += n // (5 ** exponent)
            exponent += 1
        return ans