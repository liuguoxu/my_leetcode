/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	cur, ret := 0, nums[0]

	i := 0

	//从左边找，找到一个数比它和他之前所有数之和还要大

	for i = 0; i < l; i++ {
		cur = max(cur+nums[i], nums[i])
		ret = max(cur, ret)
	}

	return ret
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
