/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/

package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	mid := 0
	found := false

	left, right := 0, len(nums)-1

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			found = true
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found {
		ret := []int{mid, mid}

		for i := mid - 1; i >= 0; i-- {
			if nums[i] == target {
				ret[0] = i
			} else {
				break
			}
		}

		for i := mid + 1; i < len(nums); i++ {
			if nums[i] == target {
				ret[1] = i
			} else {
				break
			}
		}

		return ret

	} else {
		return []int{-1, -1}
	}

}
