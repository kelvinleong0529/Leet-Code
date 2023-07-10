/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)

	var recursiveFunc func(head *ListNode)
	recursiveFunc = func(head *ListNode) {
		if head == nil {
			return
		}
		stack = append(stack, head)
		recursiveFunc(head.Next)
	}
    recursiveFunc(head)
    
    var newHead, p1 *ListNode
	for i := len(stack) - 1; i >= 0; i-- {
		if i == len(stack)-1 {
			newHead = stack[i]
			p1 = newHead
			continue
		}
		p1.Next = stack[i]
		p1 = stack[i]
		if i == 0 {
			p1.Next = nil
		}
	}
	return newHead
}
