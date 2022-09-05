class Solution:
    def firstUniqChar(self, s: str) -> int:
        hash_map = {}
        index = []
        for i,v in enumerate(s):
            if v in hash_map:
                try:
                    index.remove(hash_map[v])
                except:
                    pass
            else:
                hash_map[v] = i
                index.append(i)
        if len(index) == 0:
            return -1
        else:
            return index[0]