class Solution {
    public int[][] construct2DArray(int[] original, int m, int n) {
        if ( m*n != original.length) {
            return new int[0][0];
        }
        int[][] answer = new int[m][n];
        int row = 0;
        int col = 0;
        for (int element: original) {
            answer[row][col] = element;
            col++;
            if (col == n) {
                col = 0;
                row++;
            }
        }
        return answer;
    }
}