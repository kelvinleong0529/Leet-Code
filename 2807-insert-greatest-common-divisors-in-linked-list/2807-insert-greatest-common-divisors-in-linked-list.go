/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func greatestCommonDivisor(x, y int) int {
	var small int
	if x < y {
		small = x
	} else {
		small = y
	}
	var answer int
	for i := small; i > 0; i-- {
		if x%i == 0 && y%i == 0 {
			answer = i
			break
		}
	}
	return answer
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	answer := head
	for {
		if head == nil || head.Next == nil {
			break
		}
		slow := head
		fast := head.Next
		gcd := greatestCommonDivisor(slow.Val, fast.Val)
		gcdNode := &ListNode{
			Val:  gcd,
			Next: fast,
		}
		slow.Next = gcdNode
		head = fast
	}
	return answer
}
