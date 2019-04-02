/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?


Above is a 7 x 3 grid. How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:

Input: m = 7, n = 3
Output: 28
*/

package main

import "fmt"

func fDFS(m, n, row, column int, num *int, record map[int]map[int]int) {
	if row >= m || column >= n {
		return
	}

	if row == m-1 && column == n-1 {
		*num++
		return
	}

	if rcd, ok := record[row+1][column]; !ok {
		before := *num
		fDFS(m, n, row+1, column, num, record)
		if *num-before > 0 {
			record[row+1][column] = *num - before
		}
	} else {
		*num += rcd
	}

	if rcd, ok := record[row][column+1]; !ok {
		before := *num
		fDFS(m, n, row, column+1, num, record)
		if *num-before > 0 {
			record[row][column+1] = *num - before
		}
	} else {
		*num += rcd
	}

}

func uniquePaths(m int, n int) int {
	var record = make(map[int]map[int]int)

	for i := 0; i < m; i++ {
		record[i] = make(map[int]int)
	}

	ret := 0
	fDFS(m, n, 0, 0, &ret, record)
	return ret
}

func main() {
	ret := uniquePaths(3, 3)

	fmt.Println(ret)
}
