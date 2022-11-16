class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        hashMap = {}
        for i, v in enumerate(s):
            if v in hashMap:
                if len(hashMap[v]) == 2:
                    hashMap[v][-1] = i
                else:
                    hashMap[v].append(i)
            else:
                hashMap[v] = [i]
        ans = -1
        for _, v in hashMap.items():
            if len(v) == 2:
                if v[-1]-v[0]-1 > ans:
                    ans = v[-1]-v[0]-1
        return ans