class Solution:
    def isSumEqual(self, firstWord: str, secondWord: str, targetWord: str) -> bool:
        
        def ascii(s:str) -> str:
            return str(ord(s) - ord('a'))
        
        def getValue(s:str) -> str:
            value = ""
            for i in s:
                asc = ascii(i)
                value += asc
            return value
        
        first = getValue(firstWord)
        second = getValue(secondWord)
        target = getValue(targetWord)
        
        if int(first) + int(second) == int(target):
            return True
        else:
            return False
                    