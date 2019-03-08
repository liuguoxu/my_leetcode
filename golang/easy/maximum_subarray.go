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
	var cur int

	i := 0

	//从左边找，找到一个数比它和他之前所有数之和还要大
	lmax := -1
	for i = 0; i < len(nums); i++ {
		if nums[i] > cur+nums[i] {
			cur = nums[i]
			lmax = i
		} else {
			cur = cur + nums[i]
		}
	}

	//从右边找，找到一个数比它和他之前所有数之和还要大
	rmax := -1
	cur = 0
	for i = len(nums) - 1; i >= 0; i-- {
		if nums[i] > cur+nums[i] {
			cur = nums[i]
			rmax = i
		} else {
			cur = cur + nums[i]
		}
	}

	fmt.Println(lmax, rmax)

	ret := 0
	if lmax == -1 {
		lmax = 0
	}
	if rmax == -1 {
		rmax = len(nums) - 1
	}

	if lmax >= rmax {
		return max(nums[lmax], nums[rmax])
	}

	for k := lmax; k <= rmax; k++ {
		fmt.Println(nums[k])
		ret += nums[k]
	}
	return ret
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
