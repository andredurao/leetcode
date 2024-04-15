// https://leetcode.com/problems/diameter-of-binary-tree/description/
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

func sumNumbers(root *TreeNode) int {
	total := 0
	traverse(root, root.Val, &total)
	return total
}

func traverse(node *TreeNode, pathTotal int, total *int) {
	if node.Left == nil && node.Right == nil {
		*total += pathTotal
		return
	}
	if node.Left != nil {
		traverse(node.Left, (pathTotal*10 + node.Left.Val), total)
	}
	if node.Right != nil {
		traverse(node.Right, (pathTotal*10 + node.Right.Val), total)
	}
}

// ------------------------------------------------------

func main() {
	root := BuildTree("[2,0,0]")
	fmt.Println(sumNumbers(root))
}
