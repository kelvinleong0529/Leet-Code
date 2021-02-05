class Solution:
    def isHappy(self, n: int) -> bool:
        temp = set()
        while 1:
            ans = 0
            while n > 0:
                ans += (n%10)**2
                n = n//10
            if ans == 1:
                return 1
            elif ans in temp:
                return 0
            else:
                temp.add(ans)
                n = ans