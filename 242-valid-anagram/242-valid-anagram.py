class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        hash_table = {}
        for i in s:
            if i in hash_table:
                hash_table[i] += 1
            else:
                hash_table[i] = 1

        for i in t:
            if i not in hash_table:
                return False
            hash_table[i] -= 1
            if hash_table[i] == 0:
                del hash_table[i]
        
        return True if len(hash_table) == 0 else False