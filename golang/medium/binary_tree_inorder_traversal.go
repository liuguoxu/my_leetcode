/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)

	ret := make([]int, 0)

	if root == nil {
		return ret
	}

	p := root

	for len(stack) != 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		out := stack[len(stack)-1]
		ret = append(ret, out.Val)

		stack = stack[:len(stack)-1]

		p = out.Right
	}

	return ret
}
