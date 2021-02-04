class Solution:
    def convertToTitle(self, n: int) -> str:
        ans = ''
        while n is not -1:
            rem = (n-1) % 26 
            n = int((n-1)/26)
            ans = chr(65+rem) + ans
            if n==0:
                n = -1
        return ans