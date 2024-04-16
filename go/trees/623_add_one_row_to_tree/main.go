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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	res := root
	if depth == 1 {
		res = &TreeNode{Val: val, Left: root, Right: nil}
	} else {
		traverseAddRow(root, val, depth, 1)
	}
	return res
}

func traverseAddRow(node *TreeNode, val int, depth int, curDepth int) {
	if node == nil {
		return
	}
	if curDepth == depth-1 {
		l := node.Left
		r := node.Right
		node.Left = &TreeNode{Val: val, Left: l, Right: nil}
		node.Right = &TreeNode{Val: val, Left: nil, Right: r}
	}
	traverseAddRow(node.Left, val, depth, curDepth+1)
	traverseAddRow(node.Right, val, depth, curDepth+1)
}

// ------------------------------------------------------

func main() {
	input := "[4,2,6,3,1,5]"
	root := BuildTree(input)
	// addOneRow(root, 1, 2)
	addOneRow(root, 1, 1)
	// fmt.Println(input)
	fmt.Println(TreeToArray(root))
}
