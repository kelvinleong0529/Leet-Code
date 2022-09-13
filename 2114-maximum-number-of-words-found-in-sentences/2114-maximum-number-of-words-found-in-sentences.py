class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        ans = 0
        for sentence in sentences:
            word_count = 0
            for s in sentence:
                if s == " ":
                    word_count += 1
            if word_count + 1 > ans:
                ans = word_count + 1
        
        return ans