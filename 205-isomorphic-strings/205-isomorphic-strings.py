class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        hash_map = {}
        is_changed = {}
        for i in range(len(s)):
            s1 = s[i]
            t1 = t[i]
            if s1 in hash_map:
                if hash_map[s1] != t1:
                    return False
            else:
                if t1 in is_changed and is_changed[t1]:
                    return False
            is_changed[t1] = True
            hash_map[s1] = t1
        return True