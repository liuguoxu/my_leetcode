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

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
//ret:
// -1- duplicated
// other - length
int add_hash(char *hash, char value)
{
	printf("%c\n", value);
    int i = 0;
    for (; i < strlen(hash); i++) {
        if (hash[i] == value)
            return -1;
    }
    hash[i] = value;

	printf("%s,%ld\n", hash, strlen(hash));
    return strlen(hash);
}

int lengthOfLongestSubstring(char *s)
{
    int l = strlen(s);

    char *hash = calloc(l + 1, sizeof(char));

    int max = 0;
    int ret = 0;
    for (int i = 0; i < l; ++i) {
        if ((ret = add_hash(hash, s[i])) != -1) {
            max = ret > max? ret:max;
			continue;
		}
		i -= strlen(hash);
        memset(hash, 0, l + 1);
    }


    free(hash);
    return max;
}

int main()
{
    char *str = "mkhazkgrwfnmzelloioditgzihtersmfftfggnxctnqurdfilwltrtuzdofuirbnnjvtwtxrgnm";

    printf("%d\n", lengthOfLongestSubstring(str));
}
