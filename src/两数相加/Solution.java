package 两数相加;

import java.security.interfaces.RSAKey;

/*
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}

public class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int carry = 0;
        ListNode result = null, head = null;
        while (l1 != null && l2 != null) {
            int sum = l1.val + l2.val + carry;
            if (sum / 10 > 0) {
                carry = 1;
                sum %= 10;
            }
            else {
                carry = 0;
            }

            if (result == null) {
                result = new ListNode(sum);
                head = result;
            }
            else {
                result.next = new ListNode(sum);
                result = result.next;
            }
            l1 = l1.next;
            l2 = l2.next;
        }

        while (l1 != null) {
            int sum = l1.val + carry;
            if (sum / 10 > 0) {
                carry = 1;
                sum %= 10;
            }
            else {
                carry = 0;
            }
            if (result == null) {
                result = new ListNode(sum);
                head = result;
            }
            else {
                result.next = new ListNode(sum);
                result = result.next;
            }
            l1 = l1.next;
        }

        while (l2 != null) {
            int sum = l2.val + carry;
            if (sum / 10 > 0) {
                carry = 1;
                sum %= 10;
            }
            else {
                carry = 0;
            }
            if (result == null) {
                result = new ListNode(sum);
                head = result;
            }
            else {
                result.next = new ListNode(sum);
                result = result.next;
            }
            l2 = l2.next;
        }

        if (carry == 1) {
            result.next = new ListNode(1);
        }

        return head;
    }
<<<<<<< HEAD
=======

    private int getLengthOfList(ListNode head) {
        int length = 0;
        while (head != null) {
            length++;
            head = head.next;
        }
        return length;
    }
>>>>>>> 605f1bfeac621b5ced241fca22b2f84eda0133e6
}