class Solution:
    def candy(self, ratings: List[int]) -> int:
        base = [1]*len(ratings)
        left = [1]
        right = [1]
        
        for i in range(1, len(ratings)):
            if ratings[i] > ratings[i-1]:
                left.append(left[i-1]+1)
            else:
                left.append(1)
            if ratings[len(ratings)-i-1] > ratings[len(ratings)-i]:
                right.append(right[i-1]+1)
            else:
                right.append(1)
            
        right = right[::-1]
        ans = 0
        
        for i in range(len(left)):
            if left[i] >= right[i]:
                ans += left[i]
            else:
                ans += right[i]
        
        return ans