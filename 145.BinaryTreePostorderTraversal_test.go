/**
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。



 示例 1：


输入：root = [1,null,2,3]
输出：[3,2,1]


 示例 2：


输入：root = []
输出：[]


 示例 3：


输入：root = [1]
输出：[1]




 提示：


 树中节点的数目在范围 [0, 100] 内
 -100 <= Node.val <= 100




 进阶：递归算法很简单，你可以通过迭代算法完成吗？

 Related Topics 栈 树 深度优先搜索 二叉树 👍 1138 👎 0

*/

package cn

import (
	"testing"
)

func TestBinaryTreePostorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: nil}}
	t.Log(postorderTraversal(root))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 后续遍历: 左 右 根
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)

	}

	postorder(root)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
