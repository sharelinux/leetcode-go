/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

 有效字符串需满足：


 左括号必须用相同类型的右括号闭合。
 左括号必须以正确的顺序闭合。
 每个右括号都有一个对应的相同类型的左括号。




 示例 1：


输入：s = "()"
输出：true


 示例 2：


输入：s = "()[]{}"
输出：true


 示例 3：


输入：s = "(]"
输出：false




 提示：


 1 <= s.length <= 10⁴
 s 仅由括号 '()[]{}' 组成


 Related Topics 栈 字符串 👍 4320 👎 0

*/

package cn

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	s := "([)]"
	t.Log(isValid(s))
}

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
