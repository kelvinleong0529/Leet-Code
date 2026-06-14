/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {

	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}

	var bfs func()
	bfs = func() {
		for len(queue) > 0 {
			curQueueLen := len(queue)
			curAns := make([]int, curQueueLen)

			for i := range curQueueLen {
				node := queue[0]
				queue = queue[1:]
				curAns[i] = node.Val

				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}

			ans = append([][]int{curAns}, ans...)
		}
	}

	bfs()

	return ans
}
