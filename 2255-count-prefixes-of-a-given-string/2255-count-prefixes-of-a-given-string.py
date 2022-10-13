class Solution:
    def countPrefixes(self, words: List[str], s: str) -> int:
        ans = 0
        for word in words:
            if word[0] == s[0]:
                valid = True
                i = 0
                for i in range(len(word)):
                    if i >= len(s) or word[i] != s[i]:
                        valid = False
                        break
                if valid:
                    ans += 1
        
        return ans