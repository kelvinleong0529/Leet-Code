class Solution:
    def climbStairs(self, n: int) -> int:
        def fac(n,table):
            if n == 0:
                return 1
            if n in table:
                return table[n]
            ans = table[len(table)]
            i = len(table)+1
            while i < n+1 :
                ans *= i
                table[i] = ans
                i += 1
            return table[n]
        ones = n
        twos = 0
        ans = 0
        table = {1:1} # declare the table
        while True:
            ans += fac(ones+twos,table)/(fac(ones,table)*fac(twos,table))
            if ones>=2:
                ones -= 2
                twos += 1
            else:
                break
        return int(ans)