/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	head := new(TreeNode)
	cursor := head

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)

		cursor.Right = &TreeNode{Val:node.Val}
		cursor = cursor.Right

		inOrder(node.Right)
	}
    
    inOrder(root)

	return head.Right
}