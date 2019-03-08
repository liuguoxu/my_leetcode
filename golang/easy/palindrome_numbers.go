/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	//将数字翻转
	//比较x和reverse的大小，使得不需要遍历整个x，只需走到一半即可判定是否是回文
	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}

	fmt.Println(reverse, x)
	return reverse == x || x == reverse/10

}

func main() {
	fmt.Println(isPalindrome(1234321))
}
