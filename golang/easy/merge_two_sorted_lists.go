/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	p := l1
	q := l2

	var ret ListNode
	tail := &ret

	for p != nil || q != nil {
		if p == nil {
			tail.Next = &ListNode{Val: q.Val}
			tail = tail.Next
			q = q.Next
			continue
		}

		if q == nil {
			tail.Next = &ListNode{Val: p.Val}
			tail = tail.Next
			p = p.Next
			continue
		}

		if p.Val <= q.Val {
			tail.Next = &ListNode{Val: p.Val}
			tail = tail.Next
			p = p.Next
			continue
		} else {
			tail.Next = &ListNode{Val: q.Val}
			tail = tail.Next
			q = q.Next
			continue
		}
	}

	return ret.Next
}
