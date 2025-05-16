package 最长回文子串;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/10/10 3:38 下午
 * @description
 **/
class Solution {
    public String longestPalindrome(String s) {
        /** dp[i][j] 是 字符串的s的从i到j的子串是否为回文串
         *
         *  dp[i][i] = true
         *  i~j的距离为 0 dp[i][j] = true
         *  i~j的距离为 1 dp[i][j] = s[i] == s[j]
         *  i~j的距离> 1 dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
         *
         */

        int length = s.length();
        boolean[][] dp = new boolean[length][length];

        // 长度 - 1 小于 1  也就是 长度为1或0的字符串  不用递推 直接就是 回文
        if (length < 2) {
            return s;
        }

        String res = "";

        for (int i = length - 1; i >= 0; --i) {
            for (int j = i; j < length; ++j) {
                if (j == i) {
                    dp[i][j] = true;
                }
                else {
                    boolean equal = s.charAt(i) == s.charAt(j);
                    if (j - i == 1) {
                        dp[i][j] = equal;
                    }
                    else {
                        dp[i][j] = dp[i + 1][j - 1] && equal;
                    }
                }

                if (dp[i][j] && res.length() < (j - i + 1)) {
                    res = s.substring(i, j + 1);
                }
            }
        }

        return res;
    }
}