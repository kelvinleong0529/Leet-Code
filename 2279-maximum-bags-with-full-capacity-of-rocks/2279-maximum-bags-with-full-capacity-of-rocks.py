class Solution:
    def maximumBags(self, capacity: List[int], rocks: List[int], additionalRocks: int) -> int:
        for i in range(len(capacity)):
            capacity[i] -= rocks[i]
        capacity.sort()
        for i in range(len(capacity)):
            additionalRocks -= capacity[i]
            if additionalRocks == 0:
                return i+1
            if additionalRocks < 0:
                return i
        return len(rocks)