// @lc app=leetcode.cn id=1512 lang=golang
// 1512. 好数对的数目
// 给你一个整数数组 nums 。

// 如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。

// 返回好数对的数目。

// 示例 1：

// 输入：nums = [1,2,3,1,1,3]
// 输出：4
// 解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
// 示例 2：

// 输入：nums = [1,1,1,1]
// 输出：6
// 解释：数组中的每组数字都是好数对
// 示例 3：

// 输入：nums = [1,2,3]
// 输出：0
 

// 提示：

// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100
// 通过次数6,993提交次数7,791

// @lc code=start
func numIdenticalPairs(nums []int) int {
	count := 0
	for k1, v1 := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if k1 < k2 && v1 == nums[k2] {
				count++
			}
		}
	}
	return count
}
// @lc code=end