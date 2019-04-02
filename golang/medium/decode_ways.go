/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/

package main

import (
	"fmt"
	"strconv"
)

func fDFS(dict map[int]byte, s string, index int, ret *int) {
	if index >len(s) -1 {//index到最后
		return
	}


	num, err := strconv.Atoi(s[index : index+1])
	if err != nil {
		panic("invalid")
	}

	if _,ok := dict[num];!ok{//非法字符
		panic("invalid")
	}

	if index == len(s)-1 {//最后一个合法字符
		*ret++
		return
	}

	num_1, err := strconv.Atoi(s[index+1 : index+2])
	if err != nil {
		panic("invalid")
	}




	if num*10+num_1 == 10 || num*10+num_1 == 20{
		if index == len(s) -2 {//最后两位
			*ret++
			return
		}

		fDFS(dict, s, index+2, ret)
		return
	}

	if _, ok := dict[num*10+num_1]; ok {
		if index == len(s) -2 {//最后两位
			*ret++
			fDFS(dict, s, index+1, ret)
			return
		}

		fDFS(dict, s, index+2, ret)
	}

	fDFS(dict, s, index+1, ret)
}

func numDecodings(s string) int {
	dict := map[int]byte{
		1:  'A',
		2:  'B',
		3:  'C',
		4:  'D',
		5:  'E',
		6:  'F',
		7:  'G',
		8:  'H',
		9:  'I',
		10: 'J',
		11: 'K',
		12: 'L',
		13: 'M',
		14: 'N',
		15: 'O',
		16: 'P',
		17: 'Q',
		18: 'R',
		19: 'S',
		20: 'T',
		21: 'U',
		22: 'V',
		23: 'W',
		24: 'X',
		25: 'Y',
		26: 'Z',
	}

	if len(s) == 0 {
		return 0
	}

	s1, err := strconv.Atoi(s[0:1])
	if err != nil {
		return 0
	}

	if _,ok:= dict[s1];!ok{
		return 0
	}

	ret := 0


	defer func() {
		r := recover()
		if r!=nil{
			fmt.Println(r)
			ret = 0
		}
	}()

	fDFS(dict, s, 0, &ret)
	return ret
}

func main() {
	ret := numDecodings("110")
	fmt.Println(ret)
}
