class Solution {
    int ROWS;
    int COLS;

    public int countBattleships(char[][] board) {
        ROWS = board.length;
        if (ROWS == 0) {
            return 0;
        }
        COLS = board[0].length;
        int count = 0;
        for (int row = 0; row < ROWS; row++) {
            for (int col = 0; col < COLS; col++) {
                if (board[row][col] == 'X') {
                    count++;
                    dfs(board, row, col);
                }
            }
        }
        return count;
    }

    public void dfs(char[][] board, int row, int col) {
        if (row < 0 || row >= ROWS || col < 0 || col >= COLS || board[row][col] == '.') {
            return;
        }
        board[row][col] = '.';

        dfs(board, row-1, col);
        dfs(board, row+1, col);
        dfs(board, row, col-1);
        dfs(board, row, col+1);
    }
}