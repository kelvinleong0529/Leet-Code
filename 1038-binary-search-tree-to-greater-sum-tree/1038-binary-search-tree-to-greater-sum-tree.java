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
    public TreeNode bstToGst(TreeNode root) {
        dfs(root, new int[]{0});
        return root;
    }

    public static void dfs(TreeNode node, int[] sum) {
        if (node == null) {
            return;
        }
        dfs(node.right, sum);
        int value = node.val;
        node.val += sum[0];
        sum[0] += value;
        dfs(node.left, sum);
    }
}