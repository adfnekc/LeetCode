/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (61.15%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    54.8K
 * Total Submissions: 89.6K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */
package main

// @lc code=start
func getRow(rowIndex int) []int {
	var ret []int
	ret = append(ret, 1)
	for r := 1; r <= rowIndex; r++ {
		var row []int
		for n := 1; n <= r+1; n++ {
			if n-2 < 0 || n > len(ret) {
				row = append(row, 1)
			} else {
				row = append(row, ret[n-2]+ret[n-1])
			}
		}
		ret = row
	}
	return ret
}

// @lc code=end
