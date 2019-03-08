/*Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

package main

import (
	"fmt"
)

func isMatch(left, right string) bool {
	if right == ")" {
		return left == "("

	}
	if right == "]" {
		return left == "["

	}
	if right == "}" {
		return left == "{"
	}

	return false
}

func isValid(s string) bool {
	var stack string
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		switch c {
		case "(", "[", "{":
			stack = stack + c

		case ")", "]", "}":
			if stack == "" {
				return false
			}

			if !isMatch(string(stack[len(stack)-1]), c) {
				return false
			}

			stack = stack[0 : len(stack)-1]

		default:
			return false
		}
	}

	return stack == ""
}

func main() {
	fmt.Println(isValid("{[]}"))
}
