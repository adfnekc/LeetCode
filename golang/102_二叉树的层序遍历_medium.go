package main

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (64.22%)
 * Likes:    847
 * Dislikes: 0
 * Total Accepted:    294.2K
 * Total Submissions: 458.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层序遍历结果：
 *
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
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
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	levelOrderHelper(root, &res, 0)
	return res
}

func levelOrderHelper(root *TreeNode, resp *[][]int, depth int) {
	if root == nil {
		return
	}
	if depth == len(*resp) {
		*resp = append((*resp), []int{})
	}
	(*resp)[depth] = append((*resp)[depth], root.Val)
	levelOrderHelper(root.Left, resp, depth+1)
	levelOrderHelper(root.Right, resp, depth+1)
}

// @lc code=end
