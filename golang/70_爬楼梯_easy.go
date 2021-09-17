package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (51.81%)
 * Likes:    1618
 * Dislikes: 0
 * Total Accepted:    425.6K
 * Total Submissions: 821.1K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make(map[int]int)
	dp[1] = 1
	dp[2] = 2
	return climbStairsHelper(n, dp)
}

func climbStairsHelper(n int, dp map[int]int) int {
	if n == 0 {
		return 0
	}
	if step, ok := dp[n]; ok {
		return step
	}
	dp[n] = climbStairsHelper(n-1, dp) + climbStairsHelper(n-2, dp)
	return dp[n]
}

// @lc code=end
