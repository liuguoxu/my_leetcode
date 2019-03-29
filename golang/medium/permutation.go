/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

package main

func fDFS(nums []int, level int, dict map[int]bool, one []int, res *[][]int) {
	if level == len(nums) {
		*res = append(*res, one)
		return
	}

	//deep copy map
	my_dict := make(map[int]bool)
	for k, v := range dict {
		my_dict[k] = v
	}

	//deep copy slice
	my_one := make([]int, len(one))
	for i := 0; i < len(one); i++ {
		my_one[i] = one[i]
	}

	for i, v := range nums {
		if my_dict[i] {
			continue
		} else {
			my_dict[i] = true
			my_one = append(my_one, v)
			fDFS(nums, level+1, my_dict, my_one, res)

			my_dict[i] = false
			my_one = my_one[:len(my_one)-1]
		}
	}

}

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	dict := make(map[int]bool)

	fDFS(nums, 0, dict, []int{}, &res)

	return res
}
