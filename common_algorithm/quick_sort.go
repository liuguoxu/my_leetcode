//quick sort

package main

import (
	"fmt"
)

//分区排序
func partition(arr []int, left, right int) int {
	var i, j int
	for i, j = left, left; j <= right-1; j++ {
		if arr[j] < arr[right] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func main() {
	arr := []int{100, 324, 1, 342, 4, 6, 4, 5}

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)

}
