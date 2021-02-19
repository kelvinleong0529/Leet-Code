class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        drank = numBottles
        remain = 0
        temp = 0
        while (numBottles+remain)//numExchange:
            drank += (numBottles+remain)//numExchange
            temp = (numBottles+remain) % numExchange
            numBottles = (numBottles+remain)//numExchange
            remain = temp
        return drank