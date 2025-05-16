package 整数反转;

import java.util.LinkedList;
import java.util.PriorityQueue;
import java.util.Queue;
import java.util.Stack;

public class Solution {
    public int reverse(int x) {
        //正数溢出会变为负数 同理负数溢出会变为正数
        int result = 0;
        while (x != 0) {
            int num = x % 10;
            x = x / 10;

            if(result > Integer.MAX_VALUE / 10 || result == Integer.MAX_VALUE / 10 && num > 7){
                return 0;
            }

            if(result < Integer.MIN_VALUE / 10 || result == Integer.MIN_VALUE / 10 && num < -8){
                return 0;
            }

            result = result * 10 + num;
        }
        return result;
    }
}
