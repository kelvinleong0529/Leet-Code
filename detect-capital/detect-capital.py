class Solution:
    def detectCapitalUse(self, word: str) -> bool:
        if word.isupper() or word.islower():
            return 1
        elif word[0].isupper() and word[1:].islower():
            return 1
        else:
            return 0
