func numIslands(grid [][]byte) int {
    visited := make(map[[2]int]bool)
    island := 0

    rows, cols := len(grid), len(grid[0])
    directions := [][]int {{0,1}, {1,0}, {0,-1}, {-1,0}}

    var bfs func(int, int)
    bfs = func(row int, col int) {
        queue := [][]int {{row, col}}
        visited[[2]int {row, col}] = true
        
        for len(queue) != 0 {
            r, c := queue[0][0], queue[0][1]
            queue = queue[1:]

            for _, direction := range directions {
                dr := direction[0]
                dc := direction[1]

                newRow := r + dr
                newCol := c + dc

                if newRow >= 0 && newRow < rows && 
                newCol >= 0 && newCol < cols && 
                grid[newRow][newCol] == '1' {
                    if _, ok := visited[[2]int {newRow, newCol}]; !ok {
                        queue = append(queue, []int {newRow, newCol})
                        visited[[2]int {newRow, newCol}] = true
                    }
                }
            }
        }
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            _, ok := visited[[2]int {r, c}]
            if grid[r][c] == '1' && !ok {
                island++
                bfs(r, c)
            }
        }
    }

    return island
}