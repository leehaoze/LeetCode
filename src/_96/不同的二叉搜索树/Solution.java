package _96.不同的二叉搜索树;

public class Solution {
    public int cal(int n, int[] dict) {
        if (n == 1) {
            return 1;
        }

        if (dict[n] != 0) {
            return dict[n];
        }

        int result = 0;
        for (int i = 1; i < n - 1; i++) {
            result += cal(i, dict) * cal(n - 1 - i, dict);
        }

        result += 2 * cal(n - 1, dict);

        dict[n] = result;
        return result;
    }

    public int numTrees(int n) {
        int dict[] = new int[n + 1];
        return cal(n, dict);
    }
}
