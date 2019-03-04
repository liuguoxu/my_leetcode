//
// Created by jere on 2019-3-4.
//

/*
 * Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

//ret:
// 0- succ
// -1- duplicated
int add_hash(char *hash, char value)
{
    int i = 0;
    for (; i < strlen(hash); i++) {
        if (hash[i] == value)
            return -1;
    }
    hash[i] = value;

    return 0;
}

int lengthOfLongestSubstring(char *s)
{
    int l = strlen(s);

    char *hash = calloc(l + 1, sizeof(char));

    for (int i = 0; i < l - 1; ++i) {
        if (add_hash(hash, s[i]) == 0)
            continue;

        memset(hash, 0, l + 1);
    }

    int ret = strlen(hash);

    free(hash);
    return ret;
}