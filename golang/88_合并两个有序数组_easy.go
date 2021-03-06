/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (47.87%)
 * Likes:    525
 * Dislikes: 0
 * Total Accepted:    158.3K
 * Total Submissions: 330.6K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 *
 *
 *
 * 说明:
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 *
 */
package main

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	len := m + n - 1
	n--
	m--
	for i := len; i >= 0; i-- {
		if n < 0 && m < 0 {
			break
		}
		if n < 0 {
			nums1[i] = nums1[m]
			m--
		} else if m < 0 {
			nums1[i] = nums2[n]
			n--
		} else if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else if nums1[m] <= nums2[n] {
			nums1[i] = nums2[n]
			n--
		}
	}
	print(nums1)
}

// @lc code=end
