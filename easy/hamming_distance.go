/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
*/

package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	//异或  不同位置1
	tmp := x ^ y
	count := 0
	for ; tmp != 0; count++ {
		//减1 与，每次操作减少一位为1的数
		tmp &= (tmp - 1)
	}
	return count
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
