class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) == 1:
            return intervals
        intervals.sort()
        ans = []
        index = 0
        while True:
            start_time, end_time = intervals[index][0], intervals[index][-1]
            next_start_time = intervals[index+1][0]
            while next_start_time <= end_time:
                index += 1
                end_time = max(end_time, intervals[index][-1])
                if index == len(intervals) - 1:
                    ans.append([start_time, end_time])
                    return ans
                next_start_time = intervals[index+1][0]
            ans.append([start_time, end_time])
            index += 1
            if index == len(intervals) - 1:
                ans.append([intervals[index][0], intervals[index][-1]])
                return ans