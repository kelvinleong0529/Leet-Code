/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal (root *TreeNode) []int {
	result := make([]int, 0)

	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
        preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)

	return result
}
