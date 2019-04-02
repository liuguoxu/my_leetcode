/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

package main

import "fmt"

func subsets(nums []int) [][]int {
	l := len(nums)

	if l == 0 {
		return [][]int{}
	} else if l == 1 {
		return [][]int{{nums[0]}, {}}
	}

	sub_ret := subsets(nums[:l-1])
	ret := make([][]int, 0)

	for _, v := range sub_ret {
		ret = append(ret, v)
		nv_v := []int{nums[l-1]}

		for _, vv := range v {
			nv_v = append(nv_v, vv)
		}

		ret = append(ret, nv_v)
	}

	return ret
}

func main() {
	ret := subsets([]int{1, 2, 3, 4})

	fmt.Println(ret)
}
