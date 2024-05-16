// https://leetcode.com/problems/evaluate-boolean-binary-tree/description/
// from a previously solution written in ruby

package main

import (
	"fmt"
	. "tree_node"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	return traverse(root)
}

func traverse(node *TreeNode) bool {
	// case leaf
	if node.Val == 0 || node.Val == 1 {
		return node.Val == 1
	}

	if node.Val == 2 {
		return traverse(node.Left) || traverse(node.Right)
	}

	return traverse(node.Left) && traverse(node.Right)
}

// ------------------------------------------------------

func main() {
	input := "[2,1,3,null,null,0,1]"
	root := BuildTree(input)
	fmt.Println(evaluateTree(root))
}
