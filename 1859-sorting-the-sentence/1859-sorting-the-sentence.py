class Solution:
    def sortSentence(self, s: str) -> str:
        da = [None for i in range(9)]
        word = ""
        
        for i in s:
            if i == " ":
                continue
            if ord(i) >= ord('1') and ord(i) <= ord('9'):
                da[int(i)-1] = word
                word = ""
            else:
                word += i

        ans = ""
        
        for i in da:
            if i is None:
                break
            ans += i
            ans += " "
        
        return ans[:-1]
            