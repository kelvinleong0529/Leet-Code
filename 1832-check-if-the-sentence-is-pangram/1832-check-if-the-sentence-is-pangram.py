class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        da = [True for i in range(26)]
        
        for s in sentence:
            index = ord(s) - ord('a')
            if da[index]:
                da[index] = False
                
        for i in da:
            if i:
                return False
            
        return True