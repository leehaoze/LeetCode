package TwoSum;



public class Main {
    public static void main(String args[]){
        int []nums = {2, 7, 11, 15, 26, 40};
        int target = 55;

        Solution solution = new Solution();
        int []result = solution.twoSum(nums, target);
        for (int i = 0; i < result.length; ++i){
            System.out.println(result[i]);
        }
    }
}
