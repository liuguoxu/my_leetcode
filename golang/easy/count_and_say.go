/*The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.



Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"
*/

package main

import (
	"fmt"
)

func countAndSay(n int) string {
	ret := "1"

	for i := 1; i < n; i++ {
		var tmp string
		var cur string
		var num int
		for j := 0; j < len(ret); j++ {
			if num == 0 {
				cur = string(ret[j])
				num++
			} else {
				if cur == string(ret[j]) {
					num++
				} else {
					tmp = fmt.Sprintf("%s%d%s", tmp, num, cur)
					cur = string(ret[j])
					num = 1
				}
			}
		}

		ret = fmt.Sprintf("%s%d%s", tmp, num, cur)
	}

	return ret
}

func main() {
	fmt.Println(countAndSay(5))
}
