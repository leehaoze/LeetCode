package 最长回文子串;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/10/10 3:50 下午
 * @description
 **/
public class Main {
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.longestPalindrome("abbcbdefr"));
        System.out.println(solution.longestPalindrome("babad"));
        System.out.println(solution.longestPalindrome("cbbd"));
        System.out.println(solution.longestPalindrome("abcddcde"));
        System.out.println(solution.longestPalindrome("abcbefgab"));
        System.out.println(solution.longestPalindrome("abacab"));
        System.out.println(solution.longestPalindrome("a"));
        System.out.println(solution.longestPalindrome(""));
    }
}
