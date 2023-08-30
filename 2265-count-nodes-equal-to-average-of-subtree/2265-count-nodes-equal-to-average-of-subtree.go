/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
	var answer int
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)
		totalSum := node.Val + leftSum + rightSum
		totalCount := leftCount + rightCount + 1
		if totalSum/totalCount == node.Val {
			answer++
		}
		return totalSum, totalCount
	}
	dfs(root)
	return answer
}