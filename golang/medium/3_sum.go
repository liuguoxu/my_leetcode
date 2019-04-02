/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	l := len(nums)

	ret := make([][]int, 0)

	if l < 3 || nums[0] > 0 || nums[l-1] < 0 {
		return ret
	}

	for i := 0; i < l; i++ {
		//遇到整数就不用继续了，后面都是正数
		if nums[i] > 0 {
			break
		}

		//跳过重复的数字
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]

		x, y := i+1, l-1

		for x < y {
			if nums[x]+nums[y] == target {
				ret = append(ret, []int{nums[i], nums[x], nums[y]})

				//跳过重复的数字，寻找下一组不重复的解
				for x < y && nums[x] == nums[x+1] {
					x++
				}

				//跳过重复的数字，寻找下一组不重复的解
				for x < y && nums[y] == nums[y-1] {
					y--
				}
				x++
				y--
			} else if nums[x]+nums[y] < target {
				x++
			} else {
				y--
			}
		}
	}

	return ret
}
