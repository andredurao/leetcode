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
	traverse(root, []int{root.Val}, &total)
	return total
}

func traverse(node *TreeNode, path []int, total *int) {
	if node.Left == nil && node.Right == nil {
		pathTotal := sumPath(&path)
		*total += pathTotal
		return
	}
	if node.Left != nil {
		traverse(node.Left, append(path, node.Left.Val), total)
	}
	if node.Right != nil {
		traverse(node.Right, append(path, node.Right.Val), total)
	}
}

func sumPath(path *[]int) (total int) {
	for _, val := range *path {
		total *= 10
		total += val
	}
	return
}

// ------------------------------------------------------

func main() {
	root := BuildTree("[2,0,0]")
	fmt.Println(sumNumbers(root))
}
