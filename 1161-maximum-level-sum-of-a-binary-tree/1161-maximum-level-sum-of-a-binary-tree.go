/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	hashMap := make(map[int]int)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		hashMap[level] += node.Val
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 1)
    if len(hashMap) == 0 {
        return 1
    }
	var level, max int
	first := true
	for k, v := range hashMap {
		if first {
			level = k
			max = v
			first = false
			continue
		}
		if v > max {
			max = v
			level = k
		}
	}
	return level
}