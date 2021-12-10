package 区域和检索数组不可变303;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/11/21 9:22 下午
 * @description
 **/
class NumArray {
    //创建预处理数组，存储前缀和
    private int[] prefixSum;

    public NumArray(int[] nums) {
        //数组多申请一个长度
        prefixSum = new int[nums.length + 1];
        int sum = 0;
        //记录 prefixSum[0] = 0，防止数组越界 因为要包含第i位的元素
        prefixSum[0] = 0;
        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
            prefixSum[i + 1] = sum;
        }
    }

    public int sumRange(int i, int j) {
        return prefixSum[j + 1] - prefixSum[i];
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.sumRange(i,j);
 */