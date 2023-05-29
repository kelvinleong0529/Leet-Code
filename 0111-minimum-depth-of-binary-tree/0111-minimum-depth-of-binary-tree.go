/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {

	var dfs func(*TreeNode, int) int
	dfs = func(treeNode *TreeNode, index int) int {
		if treeNode.Left == nil && treeNode.Right == nil {
			return index
		}

		var left, right int
		var leftValid, rightValid bool
		if treeNode.Left != nil {
			leftValid = true
			left = dfs(treeNode.Left, index+1)
		}
		if treeNode.Right != nil {
			rightValid = true
			right = dfs(treeNode.Right, index+1)
		}

		if !leftValid {
			return right
		}
		if !rightValid {
			return left
		}

		if left < right {
			return left
		} else {
			return right
		}
	}

    if root == nil {
        return 0
    }
    
	return dfs(root, 1)

}