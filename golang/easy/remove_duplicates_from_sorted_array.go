/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	sp_num := -4548946
	ret := 0
	for k, v := range nums {

		//遍历之前遍历过的位置，找出重复
		// 将自身位置标记为-1
		dup := false

		//顺便找出已被标记为-1的位置，用j记录
		var i, j int
		for i, j = 0, sp_num; i < k; i++ {
			if v == nums[i] {
				nums[k] = sp_num
				dup = true
				break
			}

			if j == sp_num && nums[i] == sp_num {
				j = i
			}
		}
		if dup {
			continue
		}

		//没有重复，尝试将v放入已被标记-1的位置,并且将自身设置为-1
		//若没有则不动
		ret++
		if j != sp_num {
			nums[j] = v
			nums[k] = sp_num
		}
	}

	return ret
}

func main() {
	ii := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	ret := removeDuplicates(ii)

	fmt.Println(ii, ret)
}
