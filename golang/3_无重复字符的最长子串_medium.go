/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (36.17%)
 * Likes:    4814
 * Dislikes: 0
 * Total Accepted:    791.7K
 * Total Submissions: 2.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 * 示例 4:
 *
 *
 * 输入: s = ""
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	max := 0
	n := len(s)
	l, r := 0, 0

	for l < n {
		m := map[byte]int{}
		m[s[l]]++
		max = maxint(max, len(m))
		//fmt.Println(l, r)

		r = l + 1
		for r < n {
			if m[s[r]] != 0 {
				break
			}
			m[s[r]]++
			max = maxint(max, len(m))
			//fmt.Println(l, r)
			r++
		}
		l++
	}
	return max
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
