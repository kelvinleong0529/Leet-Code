class Solution:
    def grayCode(self, n: int) -> List[int]:
        if n == 0:
            return [0]
        ans = []
        for i in range(2**n):
            ans.append(i^(i>>1))
        return ans
            