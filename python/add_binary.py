'''
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
'''


class Solution:
    # return result and remainder
    @staticmethod
    def add(a: str, b: str, c: str) -> (str, str):
        num = int(a + b + c)

        print(num)
        if num == 0:
            return "0", "0"
        elif num == 1 or num == 10 or num == 100:
            return "1", "0"
        elif num == 11 or num == 110 or num == 101:
            return "0", "1"
        else:
            return "1", "1"

    def addBinary(self, a: str, b: str) -> str:
        remainder = "0"
        i, j = len(a) - 1, len(b) - 1

        ret = ""
        while i >= 0 or j >= 0:
            left = "0" if i < 0 else a[i]
            right = "0" if j < 0 else b[j]

            print("remainder:", remainder,"left:",left,"right:",right)
            r, remainder = self.add(left, right, remainder)
            print("r:",r)
            ret = r + ret
            print(remainder,"ret:", ret)
            i -= 1
            j -= 1

        if remainder == "1":
            ret = "1" + ret

        return ret


def main():
    print(Solution().addBinary(a="11", b="1"))



if __name__ == '__main__':
    main()
