class Solution:

    def toLower(self,s:str) -> str:
        return s.lower()
    
    def getRowIndex(self,s:str) -> int:
        hash_map = {
            1: list("qwertyuiop"),
            2: list("asdfghjkl"),
            3: list("zxcvbnm")
        }
        for key in hash_map:
            if s in hash_map[key]:
                return key
    
    def findWords(self, words: List[str]) -> List[str]:
        ans = []
        for word in words:
            row_index = None
            valid = True
            for char in word:
                lower_char = self.toLower(char)
                if row_index is None:
                    row_index = self.getRowIndex(lower_char)
                    continue
                if self.getRowIndex(lower_char) != row_index:
                    valid = False
                    break
            if valid:
                ans.append(word)
        return ans
                    