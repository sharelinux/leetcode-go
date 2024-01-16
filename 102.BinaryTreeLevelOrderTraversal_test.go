/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 1868 👎 0

*/

package cn

import (
	"testing"
)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	//Your input:[3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3, Left: &TreeNode{9, nil, nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	t.Log(levelOrder(root))
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
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		vals := make([]int, len(q))
		for i := range vals {
			node := q[0]
			q = q[1:]
			vals[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, vals)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
