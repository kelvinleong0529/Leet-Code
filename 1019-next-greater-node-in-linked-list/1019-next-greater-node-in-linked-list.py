# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def nextLargerNodes(self, head: Optional[ListNode]) -> List[int]:
        ans = []
        stack = [] # [val, index]
        index = 0
        while head:
            while stack and head.val > stack[-1][0]:
                _, stackI = stack.pop()
                ans[stackI] = head.val
            stack.append([head.val,index])
            index += 1
            head = head.next
            ans.append(0)
        
        return ans