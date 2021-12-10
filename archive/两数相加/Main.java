package 两数相加;



public class Main {
    public static void main(String args[]){
        ListNode a1 = new ListNode(7);
        ListNode a2 = new ListNode(8);
        ListNode a3 = new ListNode(9);
        a1.next = a2;
        a2.next = a3;
        ListNode b1 = new ListNode(3);
        ListNode b2 = new ListNode(2);
        b1.next = b2;


        Solution solution = new Solution();
        ListNode result = solution.addTwoNumbers(a1,b1);
        while(result != null){
            System.out.println(result.val);
            result = result.next;
        }
    }
}
