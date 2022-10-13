class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        temp = ""
        for w in word:
            temp += w
            if w == ch:
                temp = temp[::-1]
                word = temp + word[len(temp):]
                break
        
        return word