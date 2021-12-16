# 567. 字符串的排列

```
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false
 

提示：

1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

s1 = abcd s2 = eiabdfbdca

首先要弄明白这个排列是什么意思，不是说只要字符串S2包含S1中的字符即可，而是说S2中某一个连续子串完全包含S1中的字符.

原本想着先引动start指针，指到s1中第一个包含s2中字符的位置，然后扩大end指针，知道遇到一个不在s2中的字符，然后把start指针挪到end的下一个。
写代码中发现不太好写，其实可以不用这样。一开始的时候直接扩大窗口，扩大过程中