/*
Given a string S and a string T, find the minimum window in S
which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	dict := make(map[byte]int)
	count :=0

	ret := ""
	min := math.MaxInt32

	for i:=0;i<len(t);i++{
		dict[t[i]]++
	}

	fmt.Println(dict)

	for left,i:=0,0;i<len(s);i++{
		dict[s[i]]--

		//找到一个t中的字母
		if dict[s[i]]>=0{
			count++
		}

		fmt.Println("count:", count)

		//找齐了所有t中的字母
		for count == len(t){
			if i-left+1 <min{
				fmt.Println("left:",left, "i",i)
				min = i-left+1
				ret = s[left:i+1]
			}

			//开始滑动窗口左侧
			dict[s[left]]++
			if dict[s[left]]>0{//减去了一个t中的字母
				count--
			}
			left++
		}
	}

	return ret
}

func main() {
	ret := minWindow("abcdef", "ade")

	fmt.Println(ret)
}