class Solution:
    def commonFactors(self, a: int, b: int) -> int:
        if a > b:
            bigger = a
            smaller = b
        else:
            bigger = b
            smaller = a
        
        ans = 1
        for i in range(2,smaller+1,1):
            if not bigger%i and not smaller%i:
                ans += 1
        
        return ans