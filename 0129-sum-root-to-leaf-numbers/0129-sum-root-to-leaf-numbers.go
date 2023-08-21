/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var answer int
	numbers := make([]int, 0)
	var recursive func(*TreeNode, int)
	recursive = func(node *TreeNode, number int) {
        if node == nil {
			return
		}
		number = number*10 + node.Val
		if node.Left == nil && node.Right == nil {
			numbers = append(numbers, number)
			return
		}
		recursive(node.Left, number)
		recursive(node.Right, number)
	}
	recursive(root, 0)
	for _, v := range numbers {
		answer += v
	}
	return answer
}
