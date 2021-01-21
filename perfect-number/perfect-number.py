class Solution:
    def checkPerfectNumber(self, num: int) -> bool:
        temp = num - 1
        for i in range (min(2,int(num**0.5)),int(num**0.5)+1):
            if not num%i:
                temp = temp - i - int(num/i)
            if temp < 0:
                return 0
                break
        if temp == 0:
            return 1
        else:
            return 0
