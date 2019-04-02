/*
Given n non-negative integers representing the histogram's bar height where the width
of each bar is 1, find the area of largest rectangle in the histogram.

Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

The largest rectangle is shown in the shaded area, which has area = 10 unit.

Example:

Input: [2,1,5,6,2,3]
Output: 10
*/

package hard

func mmax(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func mmin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}

	ret := 0

	for i := 0; i < l; i++ {

		if i < l-1 && heights[i] <= heights[i+1] {
			continue
		} else { //i走到最后一个，或者i是局部最大值

			minH := heights[i]
			for j := i; j >= 0; j-- {
				minH = mmin(minH, heights[j])
				if minH == 0 {
					break
				}

				ret = mmax(ret, minH*(i-j+1))
			}
		}
	}

	return ret
}
