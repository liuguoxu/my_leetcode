'''
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
'''


class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0 or x == 1:
            return x

        # 平方数为顺序奇数之和
        # 1+3=4
        # 1+3+5=9
        # 1+3+5+7=16
        i, total, count = 1, 1, 0
        while total <= x:
            count += 1
            i += 2
            total += i

        return count


print(Solution().mySqrt(3))
