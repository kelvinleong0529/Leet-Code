class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        hashMap = {}
        ans = [1]*len(arr)
        for i,v in enumerate(arr):
            if -v in hashMap:
                hashMap[-v].append(i)
            else:
                hashMap[-v] = [i]

        heapq.heapify(arr)
        count = 1
        previous = None
        while arr:
            value = heapq.heappop(arr)
            if value == previous:
                continue
            previous = value
            for a in hashMap[-value]:
                ans[a] = count
            count += 1
        return ans