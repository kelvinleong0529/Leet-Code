class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        ans = ""
        x = str(x)
        for i in range(len(x)):
            ans += x[len(x)-i-1]

        while ans[0] == "0":
            if len(ans) > 1:
                ans = ans[1:]
            else:
                return 0
            
        if ans[-1] == "-":
            ans = "-" + ans[:-1]
        
        if int(ans) >2**31-1 or int(ans) < -2**31:
            return 0
        
        return ans
        