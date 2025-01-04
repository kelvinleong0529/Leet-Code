/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {

	queue := make([]*TreeNode, 1)
	ans := make([]int, 0)

	queue[0] = root
	currentLevelLength := 1
	nextLevelLength := 0

	// bread-first search
	for {

		if len(queue) == 0 || queue[0] == nil {
			break
		}

		currentNode := queue[0]
		queue = queue[1:]
    
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
			nextLevelLength++
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
			nextLevelLength++
		}
		currentLevelLength--
		if currentLevelLength == 0 {
			ans = append(ans, currentNode.Val)
			currentLevelLength = nextLevelLength
			nextLevelLength = 0
		}
	}

	return ans
}