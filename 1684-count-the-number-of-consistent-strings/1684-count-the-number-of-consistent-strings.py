class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        ans = 0
        for word in words:
            valid = True
            for w in word:
                if w not in allowed:
                    valid = False
                    break
            if valid:
                ans += 1
            
        return ans