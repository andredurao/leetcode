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
	paths := []int{}
	traverse(root, []int{root.Val}, &paths)
	total := 0
	for _, val := range paths {
		total += val
	}
	return total
}

func traverse(node *TreeNode, path []int, paths *[]int) {
	total := sumPath(&path)
	if node.Left == nil && node.Right == nil {
		fmt.Println("LEAF! ", path, total)
		*paths = append(*paths, total)
		return
	}
	if node.Left != nil {
		traverse(node.Left, append(path, node.Left.Val), paths)
	}
	if node.Right != nil {
		traverse(node.Right, append(path, node.Right.Val), paths)
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
