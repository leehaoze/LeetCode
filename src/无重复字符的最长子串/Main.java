package 无重复字符的最长子串;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/10/8 5:31 下午
 * @description
 **/
public class Main {
    public static void main(String args[]) {
        无重复字符的最长子串.Solution solution = new 无重复字符的最长子串.Solution();
        System.out.println(solution.lengthOfLongestSubstring("pwwkew"));
        System.out.println(solution.lengthOfLongestSubstring("abcabcbb"));
        System.out.println(solution.lengthOfLongestSubstring(" "));
        System.out.println(solution.lengthOfLongestSubstring("tmmzuxt"));
    }
}
