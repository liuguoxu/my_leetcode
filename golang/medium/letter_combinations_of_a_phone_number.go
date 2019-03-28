/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package main

import "fmt"

//深度优先递归
func fDFS(digits string, dict []string, level int,one string, ret *[]string) {
	if level == len(digits) {
		*ret = append(*ret, one)
		return
	}

	s := dict[digits[level] - '0']

	for _,v :=range s {
		fDFS(digits, dict, level+1, one+string([]byte{byte(v)}), ret)
	}
}


func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}

	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	ret := make([]string, 0)

	fDFS(digits, dict, 0,"", &ret)

	return ret
}

func main() {
	ret := letterCombinations("23")

	fmt.Println(ret)
}
