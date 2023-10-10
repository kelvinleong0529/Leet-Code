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
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> answer = new ArrayList<>();

        dfs(answer, root);
        return answer;
    }

    static public void dfs(List<Integer> answer, TreeNode node) {
        if (node == null) {
            return;
        }

        dfs(answer, node.left);
        answer.add(node.val);
        dfs(answer, node.right);
    }
}