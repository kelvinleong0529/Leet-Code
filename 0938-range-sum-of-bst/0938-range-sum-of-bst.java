/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int answer;
    
    public int rangeSumBST(TreeNode root, int low, int high) {
        answer = 0;

        class DFS {
            void dfs(TreeNode node) {
                if (node == null) {
                    return;
                }
                if (node.val >= low && node.val <= high) {
                    answer += node.val;
                }
                dfs(node.left);
                dfs(node.right);
            }
        }
       
        DFS dfs = new DFS();
        
        dfs.dfs(root);
        return answer;
    }
}