/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/

package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	var start bool
	var ret int
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			if !start {
				start = true
			}
			ret++
		} else {
			if start {
				return ret
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))

}
