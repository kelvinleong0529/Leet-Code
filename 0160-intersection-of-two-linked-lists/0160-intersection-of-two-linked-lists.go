/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := countLinkedListLength(headA)
	lengthB := countLinkedListLength(headB)

	newHeadA := headA
	newHeadB := headB

	differenceInLength := lengthA - lengthB
	if differenceInLength > 0 {
		for {
			if differenceInLength == 0 {
				break
			}
			newHeadA = newHeadA.Next
			differenceInLength--
		}
	} else {
		differenceInLength *= -1
		for {
			if differenceInLength == 0 {
				break
			}
			newHeadB = newHeadB.Next
			differenceInLength--
		}
	}

	for {
		if newHeadA == nil || newHeadB == nil {
			return nil
		}
		if newHeadA == newHeadB {
			return newHeadA
		}
		newHeadA = newHeadA.Next
		newHeadB = newHeadB.Next
	}
}

func countLinkedListLength(head *ListNode) int {
	length := 0
	for {
		if head == nil {
			break
		}
		length++
		head = head.Next
	}
	return length
}