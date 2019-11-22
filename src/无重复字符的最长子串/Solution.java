package 无重复字符的最长子串;

import java.util.HashMap;
import java.util.Map;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/9/26 12:37 下午
 * @description
 **/
// class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         // 记录已经找到的子串中最长的长度
//         int maxLength = 0;
//         // 字符串的长度
//         int strLength = s.length();
//         //
//         HashMap<Character, Integer> charMapToPos = new HashMap<>();
//         // 当前寻找的子串的起始位置
//         int startPos = 0;
//         while (startPos < strLength) {
//             StringBuilder temp = new StringBuilder();
//             for (int j = startPos; j < strLength; ++j) {
//                 // 发现重复
//                 if (charMapToPos.containsKey(s.charAt(j))) {
//                     if (temp.length() > maxLength) {
//                         maxLength = temp.length();
//                     }
//
//                     startPos = charMapToPos.get(s.charAt(j)) + 1;
//                     charMapToPos.clear();
//                     break;
//                 }
//                 else {
//                     temp.append(s.charAt(j));
//                     charMapToPos.put(s.charAt(j), j);
//                 }
//             }
//
//             if (temp.length() > maxLength) {
//                 maxLength = temp.length();
//                 break;
//             }
//         }
//
//         return maxLength;
//     }
// }

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int maxLength = 0;

        int leftPos = 0;

        int rightPos = 0;

        int length = s.length();

        HashMap<Character, Integer> charMapToPos = new HashMap<>();

        while (length - leftPos > maxLength && rightPos < length) {
            char c = s.charAt(rightPos);
            if (charMapToPos.containsKey(c) && charMapToPos.get(c) >= leftPos) {
                maxLength = Math.max(rightPos - leftPos, maxLength);
                leftPos = charMapToPos.get(c) + 1;
            }

            charMapToPos.put(c, rightPos);
            ++rightPos;
        }

        maxLength = Math.max(rightPos - leftPos, maxLength);

        return maxLength;
    }
}

