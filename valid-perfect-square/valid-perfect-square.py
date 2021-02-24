class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if not num**0.5%1: 
            return 1
        else:
            return 0