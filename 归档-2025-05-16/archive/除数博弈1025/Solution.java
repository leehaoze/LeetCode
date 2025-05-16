package 除数博弈1025;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/11/21 8:49 下午
 * @description
 **/
public class Solution {
    public boolean divisorGame(int N) {
        boolean[] dp = new boolean[N+1];

        if (N < 2) {
            return false;
        }

        dp[1] = false;

        for (int i = 2; i <= N; ++i) {
            dp[i] = !dp[i - 1];
        }


        return dp[N];
    }
}
