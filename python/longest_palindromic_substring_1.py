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
    def is_palindromic(self, s: str) -> bool:
        l = len(s)

        i, j = 0, l - 1

        while i < j:
            if s[i] != s[j]:
                return False

            i += 1
            j -= 1

        return True

    def longestPalindrome(self, s: str) -> str:
        ls = len(s)
        if ls <= 1:
            return s

        m = 0
        ret = ""

        for i in range(ls - 1):
            for j in range(i + 1, ls):

                if self.is_palindromic(s[i:j + 1]):
                    if j + 1 - i > m:
                        m = j + 1 - i
                        ret = s[i:j + 1]

        return ret if len(ret) > 0 else s[:1]


ret = Solution().longestPalindrome(
    "")
print(ret)
