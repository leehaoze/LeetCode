583. 两个字符串的删除操作

```
给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。

每步 可以删除任意一个字符串中的一个字符。

 

示例 1：

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
示例  2:

输入：word1 = "leetcode", word2 = "etco"
输出：4
 

提示：

1 <= word1.length, word2.length <= 500
word1 和 word2 只包含小写英文字母


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-operation-for-two-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

其实删除后的结果就是最长公共子序列,因此求出来子序列之后,拿原始长度减一下就好了

    
        