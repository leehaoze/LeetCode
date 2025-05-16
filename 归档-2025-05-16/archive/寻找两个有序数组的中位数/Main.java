package 寻找两个有序数组的中位数;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/10/8 7:58 下午
 * @description
 **/
public class Main {
    public static void main(String[] args) {
        // int[] nums1 = {1, 3};
        // int[] nums2 = {2};

        // int[] nums1 = {1, 2};
        // int[] nums2 = {3, 4};

        // int[] nums1 = {};
        // int[] nums2 = {1};

        int[] nums1 = {3};
        int[] nums2 = {-2, -1};

        Solution solution = new Solution();
        System.out.println(solution.findMedianSortedArrays(nums1, nums2));
    }
}
