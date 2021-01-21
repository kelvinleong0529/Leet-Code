class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        if k == 0:
            for i in range(len(code)):
                code[i] = 0
        if k < 0:
            code.reverse()
        if abs(k) == 1:
            code.append(code.pop(0))
        elif abs(k) > len(code)/2:
            sum = 0
            temp = []
            for i in range(len(code)):
                sum = sum + code[i]
            for i in range(len(code)):
                temp1 = sum
                for j in range(len(code)-abs(k)):
                    j = (len(code)+i-j)%len(code)
                    temp1 = temp1 - code[j]
                temp.append(temp1)
            code = temp.copy()
        else:
            temp = []
            for i in range(len(code)):
                temp1 = 0
                for j in range(abs(k)):
                    j = (j+i+1)%len(code)
                    temp1 = temp1 + code[j]
                temp.append(temp1)
            code = temp.copy()
        if k < 0:
            code.reverse()
        return code
