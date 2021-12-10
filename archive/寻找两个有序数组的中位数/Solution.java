package 寻找两个有序数组的中位数;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/10/8 7:04 下午
 * @description
 **/
class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int lengthOf1 = nums1.length;
        int lengthOf2 = nums2.length;

        boolean odd = (lengthOf1 + lengthOf2) % 2 == 0;
        int mediumPos = (lengthOf1 + lengthOf2 + 1) / 2;

        int posOf1 = 0, posOf2 = 0, currentPos = 0;

        int preNum = 0, currentNum = 0;
        double result = 0;

        while (posOf1 < lengthOf1 || posOf2 < lengthOf2) {
            preNum = currentNum;

            if (posOf1 < lengthOf1 && (posOf2 >= lengthOf2 || nums1[posOf1] < nums2[posOf2])) {
                currentNum = nums1[posOf1];
                posOf1++;
            }
            else {
                currentNum = nums2[posOf2];
                posOf2++;
            }

            currentPos++;

            // 如果是偶数
            if (odd) {
                if (currentPos == mediumPos + 1) {
                    // 到了中间位置
                    result = (preNum + currentNum) / 2D;
                }
            }
            else {
                if (currentPos == mediumPos) {
                    result = (double) currentNum;
                }
            }
        }

        return result;
    }
}
