class Solution:
    def letterCasePermutation(self, s: str) -> List[str]:
        
        ans = []
        
        def backtracking(i,string):
            if i == len(s):
                ans.append(string)
                return
            
            if ord(s[i]) >= ord("a") and ord(s[i]) <= ord("z"):
                backtracking(i+1,string+s[i].upper())
            
            if ord(s[i]) >= ord("A") and ord(s[i]) <= ord("Z"):
                backtracking(i+1,string+s[i].lower())
            
            backtracking(i+1,string+s[i])
            
        backtracking(0,"")
        
        return ans