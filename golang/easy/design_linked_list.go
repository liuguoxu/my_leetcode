/*
Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list. A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
addAtTail(val) : Append a node of value val to the last element of the linked list.
addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.
Example:

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
linkedList.get(1);            // returns 2
linkedList.deleteAtIndex(1);  // now the linked list is 1->3
linkedList.get(1);            // returns 3
Note:

All values will be in the range of [1, 1000].
The number of operations will be in the range of [1, 1000].
Please do not use the built-in LinkedList library.
*/

package main

import (
	"fmt"
)

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.next == nil {
		return -1
	}
	tmp := this
	for i := 0; i <= index; i++ {
		tmp = tmp.next
		if tmp == nil {
			return -1
		}
	}
	return tmp.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &MyLinkedList{val: val}
	tmp.next = this.next
	this.next = tmp
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.next == nil {
		this.next = &MyLinkedList{val: val}
		return
	}
	tmp := this
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &MyLinkedList{val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	//list empty
	if this.next == nil {
		if index == 0 {
			this.next = &MyLinkedList{val: val}
		}
		return
	}

	//list not empty
	tmp := this
	times := 0

	for times < index {
		times++
		tmp = tmp.next
		if tmp.next == nil {
			if times == index {
				tmp.next = &MyLinkedList{val: val}
			}
			return
		}
	}

	n := &MyLinkedList{val: val}
	n.next = tmp.next
	tmp.next = n
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.next == nil {
		return
	}

	//delete first node
	if index == 0 {
		if this.next != nil {
			this.next = this.next.next
		}
		return
	}

	tmp := this
	times := 0
	for times < index {
		times++
		tmp = tmp.next
		if tmp.next == nil {
			return
		}
	}
	tmp.next = tmp.next.next
}

func (this *MyLinkedList) PrintAll() {
	for tmp := this.next; tmp != nil; tmp = tmp.next {
		fmt.Printf("%d\t", tmp.val)
	}
	fmt.Printf("\n")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	list := Constructor()

	list.AddAtIndex(1, 2)
	list.AddAtIndex(0, 1)

}
