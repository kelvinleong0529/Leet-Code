class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        if rowIndex == 0:
            return [1]
        if rowIndex == 1:
            return [1,1]
        if rowIndex == 2:
            return [1,2,1]
        ans = [1,1]
        for i in range(3,rowIndex+2):
            temp = []
            for j in range(i):
                if j == 0 or j == i-1:
                    temp.append(1)
                else:
                    temp.append(ans[j-1]+ans[j])
            ans = temp
        return ans