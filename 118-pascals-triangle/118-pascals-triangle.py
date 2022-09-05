class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        da = [[1],[1,1]]
        if numRows <= 2:
            return da[:numRows]
        for i in range(3,numRows+1):
            temp = []
            da_temp = da[i-2]
            for j in range(i):
                if j == 0 or j == (i-1):
                    temp.append(1)
                else:
                    temp.append(da_temp[j-1]+da_temp[j])
            da.append(temp)
        return da