/*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:

Input:
    2
   / \
  1   3
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
             is 5 but its right child's value is 4.
*/

package main

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lret, rret := false, false

	if root.Val == math.MaxInt32 && root.Right != nil {
		return false
	}

	if root.Val == math.MinInt32 && root.Left != nil {
		return false
	}

	if root.Left != nil && root.Left.Val == math.MinInt32 {
		if root.Left.Left != nil {
			return false
		} else {
			lret = validate(root.Left.Right, root.Left.Val, root.Val)
		}

	} else {
		lret = validate(root.Left, math.MinInt32, root.Val)
	}

	if root.Right != nil && root.Right.Val == math.MaxInt32 {
		if root.Right.Right != nil {
			return false
		}

		rret = validate(root.Right.Left, root.Val, root.Right.Val)
	} else {
		rret = validate(root.Right, root.Val, math.MaxInt32)
	}

	return lret && rret
}
