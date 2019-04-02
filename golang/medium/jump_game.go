/*
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
*/

package main

import (
	"fmt"
	"math"
)

func canJump(nums []int) bool {

	//reach表示当前能到达的最远位置
	reach := 0
	l := len(nums)
	for i, v := range nums {
		if i > reach || reach >= l-1 {
			break
		}

		reach = int(math.Max(float64(reach), float64(i+v)))
	}

	return reach >= l-1
}

func main() {

	ret := canJump([]int{3, 2, 1, 0, 4})

	fmt.Println(ret)
}
