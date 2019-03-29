/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}

	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}

	ret := []int{}

	x, y := 0, 0

	left, right := 0, n-1
	up, down := 1, m-1

	x_now, first := true, true


	i:=0
	for i < m*n-1 {

		fmt.Println(x,y,left)
		if x_now && first {
			for y < right {
				ret = append(ret, matrix[x][y])
				y++
				i++
			}
			right--
			x_now = false
		} else if !x_now && first {
			for x < down {
				ret = append(ret, matrix[x][y])
				x++
				i++
			}
			down--
			x_now = true
			first = false

		} else if x_now && !first {
			for y > left {
				ret = append(ret, matrix[x][y])
				y--
				i++
			}
			left++
			x_now = false
		} else {
			for x > up {
				ret = append(ret, matrix[x][y])
				x--
				i++
			}
			up++
			x_now = true
			first = true
		}
	}
	ret = append(ret, matrix[x][y])

	return ret
}

func main() {
	ret := spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}})

	fmt.Println(ret)
}
