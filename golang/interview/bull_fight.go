package main

import (
	"fmt"
	"strconv"
)

const (
	BULL = 10
)

func max(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func cal_bull(str string) int {
	arr := []int{}

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 65:
			arr = append(arr, 1)

		case 74, 75, 81, 84:
			arr = append(arr, 10)

		default:
			ret, _ := strconv.Atoi(string([]byte{str[i]}))
			arr = append(arr, ret)
		}
	}

	var total int
	for _, v := range arr {
		total += v
	}
	//var ret int

	var found_bull bool
	var max_remainder int
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				tmp := arr[i] + arr[j] + arr[k]
				if tmp%10 == 0 {
					fmt.Println("found bull~", i, j, k)
					found_bull = true
					r := (total - tmp) % 10
					if r == 0 { //double bulls
						return BULL + BULL
					}
					max_remainder = max([]int{max_remainder, r})
				}
			}
		}
	}

	//one bull
	if found_bull {
		return BULL + max_remainder
	}

	return max(arr)
}

func solution(first, second string) int {
	f, s := cal_bull(first), cal_bull(second)

	fmt.Println(f, s)
	if f > s {
		return 1
	} else if f < s {
		return -1
	} else {
		return 0
	}

}

func main() {
	ret := solution("4579K", "AAAA2")

	fmt.Println(ret)
}
