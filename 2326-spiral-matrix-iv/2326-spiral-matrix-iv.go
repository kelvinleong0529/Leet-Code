/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {

	answer := make([][]int, m)
	for i := 0; i < m; i++ {
		answer[i] = make([]int, n)
		for j := 0; j < n; j++ {
			answer[i][j] = -1
		}
	}

	direction := 0
	minRow := 1
	minCol := 0
	maxRow := m - 1
	maxCol := n - 1
	row := 0
	col := 0

	var update func(int)
	update = func(i int) {
        answer[row][col] = i
        
		switch direction {
		case 0:
			col++;
			if col > maxCol {
				col--
				row++
			} else if col == maxCol {
				maxCol--
				direction = 1
			}
		case 1:
			row++
			if row == maxRow {
				maxRow--
				direction = 2
			}
		case 2:
			col--
			if col == minCol {
				minCol++
				direction = 3
			}
		case 3:
			row--
			if row == minRow {
				minRow++
				direction = 0
			}
		}
	}

	for {
		if head == nil {
			return answer
		}
		update(head.Val)
		head = head.Next
	}
}