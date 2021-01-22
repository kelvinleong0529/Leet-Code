class Solution:
    def isPalindrome(self, x: int) -> bool:
        temp = list(str(x))
        temp.reverse()
        x = list(str(x))
        return temp == x
