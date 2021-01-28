class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        count_f, count_t,ans = 0,0,1
        for i in bills:
            if i == 5:
                count_f = count_f + 1
            elif i == 10:
                if not count_f:
                    ans = 0
                    break
                count_f = count_f - 1
                count_t = count_t + 1
            else:
                if count_t and count_f:
                    count_t = count_t - 1
                    count_f = count_f - 1
                elif count_t == 0 and count_f > 2:
                    count_f = count_f - 3
                else:
                    ans = 0
                    break
        return ans