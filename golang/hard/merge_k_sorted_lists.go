/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

package hard

import "math"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	ret := &ListNode{}

	ln := len(lists)

	if ln == 0 {
		return nil
	} else if ln == 1 {
		return lists[0]
	}

	ps := make([]*ListNode, ln)

	nil_num := 0

	for i := 0; i < ln; i++ {
		ps[i] = lists[i]

		if ps[i] == nil {
			nil_num++
		}
	}


	q := ret
	for nil_num <ln{
		min :=math.MinInt32
		min_i :=-1

		for i,v := range ps{
			if v!=nil{
				min = v.Val
				min_i = i
				break
			}
		}

		for i,v := range ps{
			if v!=nil && v.Val < min{
				min = v.Val
				min_i = i

			}
		}

		q.Next = ps[min_i]
		q = q.Next
		ps[min_i] = ps[min_i].Next
		if ps[min_i] == nil{
			nil_num++
		}
	}

	return ret.Next

}
