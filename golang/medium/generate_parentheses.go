/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package main

func fDFS(left, right int, one string, ret *[]string) {

	//左括号剩余比右括号多
	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*ret = append(*ret, one)
	}

	if left > 0 {
		fDFS(left-1, right, one+"(", ret)
	}

	if right > 0 {
		fDFS(left, right-1, one+")", ret)
	}

}

func generateParenthesis(n int) []string {
	ret := make([]string, 0)

	if n == 0 {
		return ret
	}

	fDFS(n, n, "", &ret)

	return ret
}
