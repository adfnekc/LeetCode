package main

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (57.11%)
 * Likes:    436
 * Dislikes: 0
 * Total Accepted:    132.4K
 * Total Submissions: 231.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层序遍历如下：
 *
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res = [][]int{}

	zigzagLevelOrderHelper(root, &res, 0)
	return res
}

func zigzagLevelOrderHelper(root *TreeNode, resp *[][]int, depth int) {
	if root == nil {
		return
	}
	if depth == len(*resp) {
		*resp = append((*resp), []int{})
	}
	if depth%2 == 1 {
		(*resp)[depth] = append((*resp)[depth], root.Val)
	} else {
		(*resp)[depth] = append([]int{root.Val}, (*resp)[depth]...)
	}
	zigzagLevelOrderHelper(root.Right, resp, depth+1)
	zigzagLevelOrderHelper(root.Left, resp, depth+1)
}

// @lc code=end
