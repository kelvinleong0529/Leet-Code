class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        ans = ""
        min_length = len(min(strs))
        for i in range(min_length):
            temp = strs[0][i]
            for j in range(1,len(strs)):
                if strs[j][i] != temp:
                    return ans
            ans += temp
        return ans