class Solution:
    def busyStudent(self, startTime: List[int], endTime: List[int], queryTime: int) -> int:
        da = []
        for i in range(len(startTime)):
            if startTime[i] <= queryTime:
                da.append(i)
        busy = len(da)
        for i in da:
            if endTime[i] < queryTime:
                busy -= 1
        
        return busy
        