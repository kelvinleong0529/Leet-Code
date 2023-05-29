/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	sum := []int{0}
	count := []int{0}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(sum) <= level {
			sum = append(sum, 0)
			count = append(count, 0)
		}

		sum[level] += node.Val
		count[level] += 1

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	average := make([]float64, len(sum))
	for i := range sum {
		average[i] = float64(sum[i]) / float64(count[i])
	}

	return average
}