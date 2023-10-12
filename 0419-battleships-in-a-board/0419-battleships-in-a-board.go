func countBattleships(board [][]byte) int {
    rows := len(board)
    if rows == 0 {
        return 0
    }
    cols := len(board[0])

    count := 0

    // Define a DFS function to mark visited cells
    var dfs func(int, int)
    dfs = func(row, col int) {
        if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] == '.' {
            return
        }

        board[row][col] = '.' // Mark the cell as visited

        // Recursively visit neighboring cells (up, down, left, and right)
        dfs(row+1, col)
        dfs(row-1, col)
        dfs(row, col+1)
        dfs(row, col-1)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == 'X' {
                count++
                dfs(i, j) // Visit the battleship and mark its connected cells
            }
        }
    }

    return count
}
