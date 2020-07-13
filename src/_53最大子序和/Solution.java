package _53最大子序和;

class Solution {
    public int maxSubArray(int[] nums) {
        if (nums.length == 1) {
            return nums[0];
        }

        int maxValue = nums[0];
        int preSum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (preSum > 0) {
                preSum = preSum + nums[i];
            }
            else {
                preSum = nums[i];
            }

            maxValue = Math.max(preSum, maxValue);
        }

        return maxValue;

    }
}