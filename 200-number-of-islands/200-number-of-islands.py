class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid:
            return 0

        visited = set()
        island = 0

        ROW, COL = len(grid), len(grid[0])

        def bfs(row,col):
            q = collections.deque()
            visited.add((row,col))
            q.append((row,col))
            directions = [[0,1],[0,-1],[1,0],[-1,0]]

            while q:
                r,c = q.popleft()
                for dr, dc in directions:
                    R = r + dr
                    C = c + dc
                    if (R in range(ROW) and C in range(COL) and grid[R][C] == "1" and (R,C) not in visited):
                        visited.add((R,C))
                        q.append((R,C))


        for r in range(ROW):
            for c in range(COL):
                if grid[r][c] == "1" and (r,c) not in visited:
                    bfs(r,c)
                    island += 1
                
        return island