/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public int pairSum(ListNode head) {
        List<Integer> list = new ArrayList<>();
        int length = 0;

        while (head != null) {
            list.add(head.val);
            head = head.next;
            length++;
        }

        int biggestSum = 0;

        for (int i = 0; i< length/2; i++){
            int sum = list.get(i) + list.get(length-i-1);
            biggestSum = sum > biggestSum? sum : biggestSum;
        }

        return biggestSum;
    }
}