class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        ans = 0
        found = False
        for i in words:
            found = False
            temp = list(chars)
            for j in i:
                if temp.count(j):
                    temp.remove(j)
                else:
                    found = False
                    break
                found = True
            if found:
                ans += len(i)
        return ans