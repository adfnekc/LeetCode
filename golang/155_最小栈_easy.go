/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (57.62%)
 * Likes:    1139
 * Dislikes: 0
 * Total Accepted:    321.9K
 * Total Submissions: 558.6K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 *
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 *
 *
 * 提示：
 *
 *
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 *
 *
*/

package main

import "math"

// @lc code=start
type MinStack struct {
	s     []int
	min_s []int
}

func (this *MinStack) min() int {
	if len(this.min_s) == 0 {
		return math.MaxInt64
	}
	return this.min_s[len(this.min_s)-1]
}

func Constructor_MinStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	min := this.min()
	if val < min {
		min = val
	}
	this.min_s = append(this.min_s, min)
	this.s = append(this.s, val)
}

func (this *MinStack) Pop() {
	this.min_s = this.min_s[:len(this.min_s)-1]
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.min()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
