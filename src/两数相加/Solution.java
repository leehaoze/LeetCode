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
        ListNode result = new ListNode(0), head = result;
        int carry = 0;
        //这里可以一个 或 条件 只要一个不为null即可， 为null的那个赋个0就可以了
        while (l1 != null || l2 != null) {
            int x = l1 == null ? 0 : l1.val;
            int y = l2 == null ? 0 : l2.val;
            int sum = x + y + carry;

            //计算进位
            carry = sum / 10;
            sum %= 10;
            //赋值节点
            result.next = new ListNode(sum);
            result = result.next;
            //移动指针节点
            l1 = l1 == null ? null : l1.next;
            l2 = l2 == null ? null : l2.next;
        }
        //如果仍然有进位
        if (carry == 1) {
            result.next = new ListNode(1);
        }
        return head.next;
    }

    public ListNode addTwoNumbers2(ListNode l1, ListNode l2) {
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
}