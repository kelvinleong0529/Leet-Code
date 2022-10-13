class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        start, end = 0, 0
        depth = 0
        i = 0
        while i < len(s):
            if s[i] == "(":
                if depth == 0:
                    start = i
                depth += 1
            if s[i] == ")":
                depth -= 1 
                if depth == 0:
                    end = i
                    s = s[:start] + s[start+1:]
                    s = s[:end-1] + s[end:]
                    i -= 2
            i += 1
        
        return s
                    