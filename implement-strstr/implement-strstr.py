class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if not str:
            return 0
        if needle in haystack:
            return haystack.index(needle)
        else:
            return -1