/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := new(TreeNode)
	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, nums []int) {
		index := indexOfMax(nums)
		node.Val = nums[index]
		if len(nums[:index]) > 0 {
			node.Left = new(TreeNode)
			dfs(node.Left, nums[:index])
		}
		if len(nums[index+1:]) > 0 {
			node.Right = new(TreeNode)
			dfs(node.Right, nums[index+1:])
		}
	}
	dfs(root, nums)
	return root
}

func indexOfMax(slice []int) int {
	var answer int
	for i := range slice {
		if i == 0 {
			answer = i
		}
		if slice[i] > slice[answer] {
			answer = i
		}
	}
	return answer
}
