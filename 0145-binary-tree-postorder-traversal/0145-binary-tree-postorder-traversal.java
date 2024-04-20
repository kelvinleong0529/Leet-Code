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
    public List<Integer> postorderTraversal(TreeNode root) {
        List<Integer> list = new ArrayList();
        postOrderTraversal(root, list);
        return list;
    }
    
    private void postOrderTraversal(TreeNode root, List<Integer> list) {
        if (root == null) {
            return;
        }
        
        postOrderTraversal(root.left, list);
        postOrderTraversal(root.right, list);
        list.add(root.val);
    }
}