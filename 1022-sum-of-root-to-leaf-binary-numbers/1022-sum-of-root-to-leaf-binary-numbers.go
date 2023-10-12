/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = (currentSum << 1) | node.Val

	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	leftSum := dfs(node.Left, currentSum)
	rightSum := dfs(node.Right, currentSum)

	return leftSum + rightSum
}
