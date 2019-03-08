/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/
package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var ret []int

	l := len(digits)
	if l == 0 {
		return ret
	}

	i := l - 1
	digits[i] = digits[i] + 1
	for digits[i] > 9 {
		if i > 0 {
			digits[i] = digits[i] - 10
			i--
			digits[i]++
		} else {
			digits[i] = digits[i] - 10
			digits = append([]int{1}, digits...)
			return digits
		}
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
}
