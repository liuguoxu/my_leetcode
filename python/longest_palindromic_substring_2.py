'''
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
'''


class Solution:
    def find(self, s: str, left: int, right: int, start: int, m: int) -> (int, int):

        print("in, left:%d, right:%d" % (left, right))
        l = len(s)
        while left >= 0 and right <= l - 1 and s[left] == s[right]:
            left -= 1
            right += 1

        print("left:%d, right:%d" % (left, right))

        if right - left - 1 > m:
            return left + 1, right - left - 1
        else:
            return start, m

    def longestPalindrome(self, s: str) -> str:
        ls = len(s)
        if ls <= 1:
            return s

        m = 1
        left = 0

        for i in range(0, ls - 1):
            if ls - i <= m / 2:
                break

            left, m = self.find(s, i, i, left, m)
            left, m = self.find(s, i, i + 1, left, m)

        print("left:%d, m:%d" % (left, m))

        return s[left:left + m]


ret = Solution().longestPalindrome("ccc")
print(ret)
