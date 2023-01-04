package main

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (56.95%)
 * Likes:    647
 * Dislikes: 0
 * Total Accepted:    81.1K
 * Total Submissions: 142.3K
 * Testcase Example:  '10'
 *
 * 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
 *
 * 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：12
 * 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 * 解释：1 通常被视为丑数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// 2 4 6 8 10 12
// 3 6 9 12 15
//
// @lc code=start
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	var i = 1
	var i2, i3, i5 int
	var m int
	var m2 int
	for i < n {
		n2 := 2 * i2
		n3 := 3 * i3
		n5 := 5 * i5

		m = min(min(n2, n3), n5)

		if m == n2 {
			i2 = d(i2)
		} else if m == n3 {
			i3 = d(i3)
		} else {
			i5 = d(i5)
		}
		if m == m2 {
			continue
		}
		m2 = m
		i += 1
	}

	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func d(n int) int {
	n += 1
	for !(n == 1 || n%2 == 0 || n%3 == 0 || n%5 == 0) {
		n += 1
	}

	return n
}

// @lc code=end
