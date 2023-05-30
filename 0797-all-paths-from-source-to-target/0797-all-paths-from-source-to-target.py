class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        result = []
        queue = [[0]]
        target = len(graph)-1
        
        while queue:
            temp = queue.pop(0)
            if temp[-1] == target:
                result.append(temp)
            else:
                for i in graph[temp[-1]]:
                    queue.append(temp+[i])
                    
        return result