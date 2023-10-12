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
    public int countNodes(TreeNode root) {
        int[] answer = {0};
        dfs(root, answer);
        return answer[0];
    }

    public void dfs(TreeNode node, int[] count) {
        if (node == null) return;
        count[0]++;
        dfs(node.left, count);
        dfs(node.right, count);
    }
}