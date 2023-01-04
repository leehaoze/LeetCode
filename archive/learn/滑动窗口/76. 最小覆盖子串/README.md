# 76. 最小覆盖子串
```
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

 

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
 

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
 

提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

很容易想到的是使用滑动窗口来做,重点是怎么进行窗口的调整.
只要当前窗口没有包括到t中所有的字符,那么就需要扩大窗口
一个这样的case ADOBECODEBANC
理想的窗口变化是这样的
A -> AD -> ADO -> ADOB -> ADOBE -> ADOBEC -> DOBEC -> OBEC -> BEC -> BECO -> BECOD -> BECODE -> BECODEB -> ECODEB -> CODEB -> CODEBA -> ODEBA -> DEBA ->
EBA -> BA -> BAN -> BANC

之前就想到用一个map记录下,key是t中的字符,value是t中的对应字符出现的次数,当进行扩容时,发现在map中存在,就对其值进行-1,用另外一个count来记录当前窗口包含t中字符的数量.
另外对于额外的重复出现的字符,会使得map中对应的value为负数,这个用另外一个变量记录下来

那么对应的窗口缩容的条件就是 count == len(t) || repeatCount > 0 || window[0] not in map

测试的过程发现对窗口进行缩容时会有 窗口长度变为0的情况 因此内部判断要特殊处理下这种情况

测试又发现这个case出错 aaaaaaaaaaaabbbbbcdd
原因是字符b的数量超级多,导致窗口缩容将开头的a给缩掉了,因此这个窗口缩容条件改为
count == len(t) || (repeatCount > 0 && windos[0] not in map)