/*
Given two integers dividend and divisor, divide two integers without using multiplication,
division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which
could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
*/

package main

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}

	if dividend == 0 {
		return 0
	}

	flag := 0
	if (dividend > 0 && divisor > 0) ||
		(dividend < 0 && divisor < 0) {
		flag = 1
	} else {
		flag = -1
	}

	var ret int

	m, n := int(math.Abs(float64(dividend))), int(math.Abs(float64((divisor))))

	for m >= n {
		t := n
		p := 1
		if m > (t << 1) {
			t <<= 1
			p <<= 1
		}

		ret = ret + p
		m -= t
	}

	return flag * ret
}
