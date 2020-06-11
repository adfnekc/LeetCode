/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (66.58%)
 * Likes:    310
 * Dislikes: 0
 * Total Accepted:    82.7K
 * Total Submissions: 124.2K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 * 
 * 
 * 
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 * 
 * 示例:
 * 
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 */
//* [
//   * ⁠    [1],
//   * ⁠   [1,1],
//  * ⁠  [1,2,1],
//  * ⁠	[1,3,3,1],
//  [1, 4, 6, 4, 1]
//[1, 5, 10,10 ,5, 1]
//  * ]
// @lc code=start
func generate(numRows int) [][]int {
	var ret [][]int
	if numRows == 0{
		return ret
	}
	ret = append(ret,[]int{1})
	for r := 1; r < numRows; r++ {
		var row []int
		for n := 1; n <= r+1; n++ {
			fmt.Println(n)
			if n-2 < 0 || n > len(ret[r-1]){
				row = append(row,1)
			} else {
				row = append(row,ret[r-1][n-2] + ret[r-1][n-1])
			}
		}
		ret = append(ret,row)
	}
	return ret
}
// @lc code=end

