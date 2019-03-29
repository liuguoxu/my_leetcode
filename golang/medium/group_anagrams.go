/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)

	for _,v:= range strs{
		b := []byte(v)

		//单词按字母排序后插入map
		sort.Slice(b, func(i, j int) bool {
			return b[i] <b[j]
		})

		dict[string(b)] = append(dict[string(b)], v)
	}


	ret := make([][]string, 0)

	for _,v :=range dict{
		ret = append(ret, v)
	}
	return ret

}

func main() {
	ret := groupAnagrams([]string{"","",""})

	fmt.Println(ret)
}
