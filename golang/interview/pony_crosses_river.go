package main

import (
	"fmt"
)

var (
	MAX_X  int
	MAX_Y  int
	Matrix [][]int
)

//判断是否走出了棋盘范围
func isValid(yx []int) bool {
	if yx[0] > -1 && yx[0] < MAX_Y &&
		yx[1] > -1 && yx[1] < MAX_X {
		return true
	}

	return false
}

var result []int

func move(total int, yx []int) {
	//y坐标到终点，记录总和
	if isValid(yx) && yx[0] == MAX_Y-1 {
		total += Matrix[yx[0]][yx[1]]
		fmt.Println(yx, total)
		result = append(result, total)
		return
	}

	if !isValid(yx) {
		return
	}

	total += Matrix[yx[0]][yx[1]]

	//可能的走法
	//若能倒着跳，则y坐标可以-1或-2
	for _, v := range [][]int{
		{yx[1] - 1, yx[0] + 2},
		{yx[1] + 1, yx[0] + 2},
		{yx[1] - 2, yx[0] + 1},
		{yx[1] + 2, yx[0] + 1},
	} {
		move(total, v)
	}
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func solution(matrix [][]int) int {
	MAX_Y, MAX_X = len(matrix), len(matrix[0])
	Matrix = matrix

	rslt := []int{}

	for i := 0; i < MAX_X; i++ {
		move(0, []int{0, i})
		rslt = append(rslt, max(result))
	}

	return max(rslt)
}

func main() {
	ret := solution([][]int{
		{1, 2, 3},
		{2, 3, 4},
	})

	fmt.Println(ret)
}
