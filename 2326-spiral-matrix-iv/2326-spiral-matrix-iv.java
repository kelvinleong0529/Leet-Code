/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
    class Solution {

        private int[][] array;
        private int row = 0;
        private int col = 0;
        private int minRow = 1;
        private int maxRow;
        private int minCol = 0;
        private int maxCol;
        private int direction = 0;


        public int[][] spiralMatrix(int m, int n, ListNode head) {
            array = new int[m][n];
            maxRow = m-1;
            maxCol = n-1;
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    array[i][j] = -1;
                }
            }

            ListNode node = head;
            while (node != null) {
                update(node.val);
                node = node.next;
            }
            return array;
        }

        private void update(int element) {
            array[row][col] = element;

            switch (direction) {
                case 0:
                    col++;
                    if (col > maxCol) {
                        col--;
                        row++;
                    } else if (col == maxCol) {
                        maxCol--;
                        direction = 1;
                    }
                    break;
                case 1:
                    row++;
                    if (row == maxRow) {
                        maxRow--;
                        direction = 2;
                    }
                    break;
                case 2:
                    col--;
                    if (col == minCol) {
                        minCol++;
                        direction = 3;
                    }
                    break;
                case 3:
                    row--;
                    if (row == minRow) {
                        minRow++;
                        direction = 0;
                    }
                    break;
            }
        }
    }