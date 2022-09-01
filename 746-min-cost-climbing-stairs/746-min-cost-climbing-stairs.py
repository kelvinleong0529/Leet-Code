class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        da = []
        for i in range(len(cost)):
            if i == 0 or i == 1:
                da.append(cost[len(cost)-i-1])
                continue
            da.append(min(cost[len(cost)-i-1]+da[i-1],cost[len(cost)-i-1]+da[i-2]))
        return min(da[-2],da[-1])
            