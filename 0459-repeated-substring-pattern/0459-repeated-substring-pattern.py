class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        s = list(s)
        string = []
        for i in range(len(s)):
            if i == 0:
                string.append(s[i])
                continue
            if s[i] == string[0]:
                valid = True
                for j in range(i,len(s),len(string)):
                    if s[j:j+len(string)] != string:
                        valid = False
                if valid:
                    return True
            string.append(s[i])
        
        return False