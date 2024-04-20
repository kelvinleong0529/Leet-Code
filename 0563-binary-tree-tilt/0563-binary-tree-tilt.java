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
    private int answer = 0;
    
    public int findTilt(TreeNode root) {
        tilt(root);
        return answer;
    }

    private int tilt(TreeNode node) {
        if (node == null) {
            return 0;
        }
        
        int leftTilt = tilt(node.left);
        int rightTilt = tilt(node.right);
        
        answer += Math.abs(leftTilt - rightTilt);
        
        return leftTilt + rightTilt + node.val;
    }
}