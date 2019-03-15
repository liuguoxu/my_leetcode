package main

import (
	"fmt"
)

func solution(a, b []int) []int {
	la, lb := len(a), len(b)

	if la == 0 || lb == 0 {
		return []int{}
	}

	ret := []int{}

	i, j := 0, 0
	for i < la && j < lb {
		if a[i] == b[j] {
			ret = append(ret, a[i])
			i++
			j++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}

	return ret
}

func main() {
	ret := solution([]int{1, 2, 3, 4, 5, 6}, []int{2, 3, 5})

	fmt.Println(ret)
}
