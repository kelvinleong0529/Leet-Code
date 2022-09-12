class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        hash_s = {}
        hash_t = {}
        
        for i in range(len(t)):
            if i != len(s):
                s_char = s[i]
                if s_char in hash_s:
                    hash_s[s_char] += 1
                else:
                    hash_s[s_char] = 1
            t_char = t[i]
            if t_char in hash_t:
                hash_t[t_char] += 1
            else:
                hash_t[t_char] = 1
        
        for k,v in hash_t.items():
            if k not in hash_s or hash_s[k] != v:
                return k