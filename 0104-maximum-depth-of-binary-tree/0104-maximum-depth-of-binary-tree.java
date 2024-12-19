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
    public int maxDepth(TreeNode root) {
        return recursive(root);
    }

    private int recursive(TreeNode node) {
        if (node == null) {
            return 0;
        }
        int left = recursive(node.left);
        int right = recursive(node.right);
        return left > right ? left + 1 : right + 1;
    }
}