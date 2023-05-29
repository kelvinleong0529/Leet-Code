/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	var deepestLevel, sum int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		
        if node == nil {
            return
        }
        
        if level == deepestLevel {
			sum += node.Val
		}
        
        if level > deepestLevel {
			deepestLevel = level
			sum = node.Val
		}

		
		dfs(node.Left, level+1)
		dfs(node.Right, level +1)
	}

	dfs(root, 0)

	return sum
}