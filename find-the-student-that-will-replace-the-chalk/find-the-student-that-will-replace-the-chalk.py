import bisect

class Solution:
    def chalkReplacer(self, chalk: List[int], k: int) -> int:
        temp = []
        temp.append(chalk[0])
        for i in range(1,len(chalk)):
            temp.append(temp[i-1]+chalk[i])
        k = k % sum(chalk)
        return bisect.bisect(temp,k)