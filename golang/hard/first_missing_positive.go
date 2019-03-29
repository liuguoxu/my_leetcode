/*
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
*/

func firstMissingPositive(nums []int) int {
	dict := make(map[int]bool)

	max_num := -1

	for _, v := range nums {
		if v > 0 {
			dict[v] = true

			if v > max_num {
				max_num = v
			}
		}
	}

	if max_num < 0 {
		return 1
	}

	i := 1
	for ; i <= max_num; i++ {
		if !dict[i] {
			return i
		}
	}

	return i

}