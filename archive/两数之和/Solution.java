package 两数之和;

import java.util.HashMap;

public class Solution {
    public int[] twoSum(int[] nums, int target) {
        //key 为值，value为下标
        HashMap<Integer, Integer> hashMap = new HashMap<>();
        for (int i = nums.length - 1; i >= 0; --i) {
            if (hashMap.containsKey(target - nums[i])) {
                return new int[]{i, hashMap.get(target - nums[i])};
            }
            hashMap.put(nums[i], i);
        }
        return new int[]{};
    }
}