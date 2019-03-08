/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	haylen, neelen := len(haystack), len(needle)

	if neelen > haylen {
		return -1
	}

	for i := 0; i <= haylen-neelen; i++ {
		if haystack[i:neelen+i] == needle {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
