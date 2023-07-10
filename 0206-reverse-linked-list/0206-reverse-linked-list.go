/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var p1, p2 *ListNode
	for {
		if head == nil {
			break
		}
		p2 = p1
		p1 = head
        head = head.Next
		p1.Next = p2
	}
	return p1
}