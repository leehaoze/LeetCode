package 跳水板;

public class Solution {
    public int[] divingBoard(int shorter, int longer, int k) {
        if (k == 0) {
            return new int[0];
        }

        if (shorter == longer) {
            return new int[] {k * shorter};
        }

        int[] result = new int[k + 1];
        for (int i = 0; i <= k; i++) {
            result[i] = (k - i) * shorter + (i) * longer;
        }

        return result;
    }
}