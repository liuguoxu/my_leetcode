/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p,prev := head,head

	for i:=0;i<n;i++{
		p = p.Next
	}

	if p == nil{
		return head.Next
	}

	for p.Next!=nil{
		p = p.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next

	return head
}

func main() {
	ret := removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, 2)

	fmt.Println(ret)
}
