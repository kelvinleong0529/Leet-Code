class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        def fac(n,table):
            if n == 0:
                return 1
            if n in table:
                return table[n]
            i = len(table) + 1
            while i < n+1:
                table[i] = table[i-1]*i
                i += 1
            return table[n]
                
        k -= 1
        remaining = [i+1 for i in range(n)]
        size = n-1
        ans = []
        table = {1:1}
        for i in range(n):
            fac_rslt = fac(size,table)
            index = k // fac_rslt
            ans.append(str(remaining[index]))
            remaining.pop(index)
            k %= fac_rslt
            size -= 1
        return ''.join(ans)