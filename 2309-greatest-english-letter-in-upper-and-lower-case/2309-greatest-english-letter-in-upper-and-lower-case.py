class Solution:
    def greatestLetter(self, s: str) -> str:
        hash_map = {}
        for i in range(0,-26,-1):
            hash_map[i] = [False,False]
        
        for i in s:
            if i.islower():
                key = ord(i) - ord('z')
                hash_map[key][0] = True
            else:
                key = ord(i) - ord('Z')
                hash_map[key][-1] = True
            
        for key in hash_map:
            if hash_map[key][0] and hash_map[key][-1]:
                return chr(key + ord('Z'))
        
        return ""