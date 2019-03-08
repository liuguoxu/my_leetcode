/*
Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.



Example 1:

Input: A = "ab", B = "ba"
Output: true
Example 2:

Input: A = "ab", B = "ab"
Output: false
Example 3:

Input: A = "aa", B = "aa"
Output: true
Example 4:

Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true
Example 5:

Input: A = "", B = "aa"
Output: false


Note:

0 <= A.length <= 20000
0 <= B.length <= 20000
A and B consist only of lowercase letters.
*/

package main

import (
	"fmt"
)

func buddyStrings(A string, B string) bool {
	la, lb := len(A), len(B)

	if la == 0 || lb == 0 ||
		la == 1 || lb == 1 ||
		la != lb {
		return false
	}

	if A == B {
		for i := 0; i < la-1; i++ {
			for j := i + 1; j < la; j++ {
				if A[i] == A[j] {
					return true
				}
			}
		}

		return false
	} else {
		var first, second, num int

		for i := 0; i < la; i++ {
			if A[i] != B[i] {
				if num == 0 {
					first = i
					num++
				} else if num == 1 {
					second = i
					num++
				} else {
					return false
				}
			}
		}

		return A[first] == B[second] && A[second] == B[first]
	}
}

func main() {
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
}
