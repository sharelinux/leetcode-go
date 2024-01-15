/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



 示例 1:


输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


 示例 2:


输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


 示例 3:


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。




 提示：


 0 <= s.length <= 5 * 10⁴
 s 由英文字母、数字、符号和空格组成


 Related Topics 哈希表 字符串 滑动窗口 👍 9909 👎 0

*/

package cn

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		max   int = 1
		left  int = 0
		right int = 1
		m         = make(map[byte]int)
	)

	m[s[left]] = left
	for ; right < len(s); right++ {
		if _, ok := m[s[right]]; ok {
			if m[s[right]] >= left {
				left = m[s[right]] + 1
			}
		}
		m[s[right]] = right

		if right-left+1 > max {
			max = right - left + 1
		}

	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)
