/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	if num == 0 {
		return ""
	}

	if num == 1 {
		return strs[0]
	}

	var ret string

	min_len := len(strs[0])

	for _, v := range strs {
		l := len(v)
		if l < min_len {
			min_len = l
		}
	}

Loop:
	for i := 0; i < min_len; i++ {
		s := string(strs[0][i])

		for j := 1; j < num; j++ {
			if s != string(strs[j][i]) {
				break Loop
			}
		}

		ret = ret + s
	}

	return ret
}

func main() {
	ret := longestCommonPrefix([]string{"flower", "flow", "flight"})

	fmt.Println(ret)
}
