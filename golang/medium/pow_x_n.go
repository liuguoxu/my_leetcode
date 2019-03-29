/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/

package main

import (
	"fmt"
	//"math"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	//计算一半
	half := myPow(x, n/2)
	if n%2 == 0 { //n 为偶数
		return half * half
	}

	//n为奇数还需而外乘一次
	if n > 0 {
		return half * half * x
	} else {
		return half * half / x
	}

}

func main() {
	//ret := myPow(0.00001,2147483647)

	fmt.Println(0.000001 * 0.000001)
}
