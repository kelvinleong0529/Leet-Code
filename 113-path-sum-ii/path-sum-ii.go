/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)

	var dfs func(*TreeNode, *[]int, *int)

	dfs = func(node *TreeNode, leaves *[]int, sum *int) {
		if node == nil {
			return
		}

		*leaves = append(*leaves, node.Val)
		*sum += node.Val

		if node.Left == nil && node.Right == nil {
			if *sum == targetSum {
				temp := make([]int, len(*leaves))
				copy(temp, *leaves)
				ans = append(ans, temp)
			}
		}

		dfs(node.Left, leaves, sum)
		dfs(node.Right, leaves, sum)        
		*leaves = (*leaves)[:len(*leaves)-1]
		*sum -= node.Val
	}

	leaves := make([]int, 0)
    sum := 0
	dfs(root, &leaves, &sum)

	return ans
}