class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        
        hash_table = {
            "2" : "abc",
            "3" : "def",
            "4" : "ghi",
            "5" : "jkl",
            "6" : "mno",
            "7" : "pqrs",
            "8" : "tuv",
            "9" : "wxyz"
        }
        
        ans = []
        
        def backtrack(i,string):
            if i == len(digits):
                ans.append(string)
                return
            
            chars = hash_table[digits[i]]
            
            for c in chars:
                backtrack(i+1,string + c)
            
        if len(digits) == 0:
            return []
        else:
            backtrack(0,"")
            return ans