/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	previous := dummyNode
	
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			previous.Next = head.Next
		} else {
			previous = previous.Next
		}
		head = head.Next
	}
	
	return dummyNode.Next
}