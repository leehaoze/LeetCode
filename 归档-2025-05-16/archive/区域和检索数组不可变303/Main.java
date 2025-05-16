package 区域和检索数组不可变303;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/11/21 9:25 下午
 * @description
 **/
public class Main {
    public static void main(String[] args) {
        int[] nums = new int[]{-2, 0, 3, -5, 2, -1};
        NumArray numArray = new NumArray(nums);
        System.out.println(numArray.sumRange(0, 2));
        System.out.println(numArray.sumRange(2, 5));
        System.out.println(numArray.sumRange(0, 5));
    }
}
