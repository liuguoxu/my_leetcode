/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i, v := range nums2 {
			nums1[i] = v
		}
		return
	} else if n == 0 {
		return
	}

	zero := m

	for i, j := 0, 0; j < n; j++ {
		for nums1[i] <= nums2[j] && i < zero {
			i++
		}

		if i == zero { //nums1到末尾
			nums1[i] = nums2[j]
			zero++
		} else { //没到末尾，nums1搬移

			for k := zero; k > i; k-- {
				nums1[k] = nums1[k-1]
			}
			zero++
			nums1[i] = nums2[j]
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)
}
