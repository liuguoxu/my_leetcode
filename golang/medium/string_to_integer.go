package medium

/*
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/

import (
	"fmt"
	"math"
)

func deletePrefix(str string, prefix byte) string {

	for i := 0; i < len(str); i++ {
		if str[i] != prefix {
			return str[i:]
		}
	}

	return str
}

func myAtoi(str string) int {
	sl := len(str)

	if sl == 0 {
		return 0
	}

	//discard whitespaces
	nstr := deletePrefix(str, ' ')

	if len(nstr) == 0 {
		return 0
	}

	minus := false

	if nstr[0] == '-' {
		//排除前导0
		nstr = nstr[1:]

		nstr = deletePrefix(nstr, '0')
		if len(nstr) == 0 {
			return 0
		}

		minus = true
	} else if nstr[0] == '+' {
		nstr = nstr[1:]

		nstr = deletePrefix(nstr, '0')
		if len(nstr) == 0 {
			return 0
		}

	} else if nstr[0] >= '0' && nstr[0] <= '9' {
		//删除前导0
		nstr = deletePrefix(nstr, '0')
		if len(nstr) == 0 {
			return 0
		} else if nstr[0] < '0' || nstr[0] > '9' {
			fmt.Println("invalid str3")
			return 0
		}

	} else {
		fmt.Println("invalid str4!")
		return 0
	}

	ret := 0

	fmt.Println(minus)

	for i := 0; i < len(nstr) && nstr[i] >= '0' && nstr[i] <= '9'; i++ {

		num := nstr[i] - 48


		if !minus && ret >= math.MaxInt32/10 {
			if ret == math.MaxInt32/10 && num <= 7 {
				return ret*10 + int(num)
			} else {
				fmt.Println("str is too big")
				return math.MaxInt32
			}
		} else if minus && -ret <= math.MinInt32/10 {
			if -ret == math.MinInt32/10 && num <= 8 {
				return -(ret*10 + int(num))
			} else {
				fmt.Println("str is too small")
				return math.MinInt32
			}
		}

		ret = ret*10 + int(num)

	}

	if minus {
		return -ret
	}

	return ret
}

func main() {
	ret := myAtoi("-2147483649")

	fmt.Println("ret:", ret)
}
