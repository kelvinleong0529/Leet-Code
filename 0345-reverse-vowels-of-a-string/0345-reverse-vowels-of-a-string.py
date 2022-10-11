class Solution:
    def reverseVowels(self, s: str) -> str:
        
        def isVowel(s:str) -> bool:
            vowels = list("aeiouAEIOU")
            if s in vowels:
                return True
            else:
                return False
        
        def swap(s:List[str],l:int,r:int) -> List[str]:
            temp = s[l]
            s[l] = s[r]
            s[r] = temp
            return s
            
        s = list(s)
        left, right = 0, len(s)-1
                
        while left < right:
            while left < len(s) and not isVowel(s[left]):
                left += 1
            while right >= 0 and not isVowel(s[right]):
                right -= 1
            if left < len(s) and right >= 0 and left < right:
                s = swap(s,left,right)
                left += 1
                right -= 1
        
        return "".join(s)