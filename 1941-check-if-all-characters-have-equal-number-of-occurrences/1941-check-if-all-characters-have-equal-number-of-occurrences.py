class Solution:
    def areOccurrencesEqual(self, s: str) -> bool:
        hash_map = {}
        for i in s:
            if i in hash_map:
                hash_map[i] += 1
            else:
                hash_map[i] = 1
        
        previous = None
        for key in hash_map:
            if previous is None:
                previous = hash_map[key]
                continue
            if hash_map[key] != previous:
                return False
        
        return True