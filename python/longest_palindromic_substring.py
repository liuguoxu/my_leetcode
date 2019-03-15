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
    def is_palindromic(self, s:str)->bool:
        pass

    def longestPalindrome(self, s: str) -> str:
        ls = len(s)
        if ls <= 1:
            return s

        index= leng = 0

        for i in range(ls-1):
            for j in range(i+1,ls):
                if self.is_palindromic(str[i:j]):

