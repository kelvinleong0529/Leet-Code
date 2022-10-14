class Solution:
    def truncateSentence(self, s: str, k: int) -> str:
        index, spaces = 0, 0
        for i in s:
            if spaces == k:
                return s[:index-1]
            if i == " ":
                spaces += 1
            index += 1
        
        return s