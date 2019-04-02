/*
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:

Input:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
Example 2:

Input:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

package main

//使用两个map记录为0的行和列，空间复杂度为O(m+n)
func setZeroes(matrix [][]int) {

	xsl := make(map[int]bool)
	ysl := make(map[int]bool)

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				xsl[i] = true
				ysl[j] = true
			}
		}
	}

	for k := range xsl {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}

	for k := range ysl {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}

}

//使用第0行和第0列来标记哪些行和列需要被置零，无需而外空间
func new_setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	//先记录第0行第0列本身是否需要置零
	rowZero, columnZero := false, false
	rows, columns := len(matrix), len(matrix[0])

	for i := 0; i < columns; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			columnZero = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowZero {
		for i := 0; i < columns; i++ {
			matrix[0][i] = 0
		}
	}

	if columnZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}
