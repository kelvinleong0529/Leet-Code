class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        arr = [".-","-...","-.-.","-..",".","..-.","--.","....",
               "..",".---","-.-",".-..","--","-.","---",".--.",
               "--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
        
        hash_table = {}
        for word in words:
            code = ""
            for w in word:
                index = ord(w) - 97
                code += arr[index]
            if code in hash_table:
                hash_table[code] += 1
            else:
                hash_table[code] = 1
        
        return len(hash_table)