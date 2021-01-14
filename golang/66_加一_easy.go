/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (44.28%)
 * Likes:    480
 * Dislikes: 0
 * Total Accepted:    162.4K
 * Total Submissions: 366.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */
package main

// @lc code=start
func plusOne(digits []int) []int {
	l := len(digits) - 1
	if digits[l] >= 9 {
		if l == 0 {
			return []int{1, 0}
		}
		return append(plusOne(digits[:l]), 0)
	}
	digits[l] += 1
	return digits
}

// @lc code=end
