/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)

		result = append(result, node.Val)

		inOrder(node.Right)
	}

	inOrder(root)

	return result
}