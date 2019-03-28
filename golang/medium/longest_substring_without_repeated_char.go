/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package medium

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	mm := make(map[byte]int)

	var max int
	for i := 0; i < len(s); i++ {
		if in_v, ok := mm[s[i]]; ok {
			l := len(mm)

			if l > max {
				max = l
			}
			fmt.Println(i, in_v)

			i = in_v
			mm = make(map[byte]int)
			continue
		}
		mm[s[i]] = i
		l := len(mm)

		if l > max {
			max = l
		}
	}

	return max
}

func main() {
	str := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(str))
}
