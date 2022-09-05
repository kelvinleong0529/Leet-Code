class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        for i in t:
            if len(s) == 0:
                return True
            temp = s[0]
            if temp == i:
                s = s[1:]
        if len(s) == 0:
            return True
        else:
            return False
            