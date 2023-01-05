package main

/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 *
 * https://leetcode-cn.com/problems/n-th-tribonacci-number/description/
 *
 * algorithms
 * Easy (60.89%)
 * Likes:    233
 * Dislikes: 0
 * Total Accepted:    141.6K
 * Total Submissions: 232.3K
 * Testcase Example:  '4'
 *
 *泰波那契序列 Tn 定义如下：
 *
 *T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
 *
 *给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
 *
 *
 *
 *示例 1：
 *
 * 输入：n = 4
 * 输出：4
 * 解释：
 * T_3 = 0 + 1 + 1 = 2
 *T_4 = 1 + 1 + 2 = 4
 *
 *
 *示例 2：
 *
 * 输入：n = 25
 *输出：1389537
 *
 *
 *
 *
 *提示：
 *
 *
 * 0 <= n <= 37
 *答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
 *
 *
 */

// @lc code=start
func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n < 2 {
		return 1
	}
	p := 0
	q := 1
	r := 1
	tmpq := 0

	for i := 2; i < n; i++ {
		tmpq = q
		r, q = p+q+r, r
		p = tmpq
	}
	return r
}

// @lc code=end
