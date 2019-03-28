/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines,
which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case,
the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	var ret int

	i, j := 0, len(height)-1

	for i < j {

		fmt.Println(i,j)

		arr := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)

		ret = int(math.Max(float64(arr), float64(ret)))

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return ret

}

func main() {
	maxArea([]int{2,3,4,5,18,17,6})
}
