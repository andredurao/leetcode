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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	tmp := root
	traverseAndDelete(nil, root, -1, &target)

	// single node with target value case
	if tmp != nil && tmp.Left == nil && tmp.Right == nil && tmp.Val == target {
		return nil
	}
	return tmp
}

func traverseAndDelete(prev *TreeNode, cur *TreeNode, dir int, target *int) {
	if cur == nil {
		return
	}

	traverseAndDelete(cur, cur.Left, 0, target)
	traverseAndDelete(cur, cur.Right, 1, target)

	if prev != nil && cur.Left == nil && cur.Right == nil && cur.Val == *target {
		if dir == 0 { // Left
			prev.Left = nil
		} else if dir == 1 {
			prev.Right = nil
		}
		return
	}
}

// ------------------------------------------------------

func main() {
	input := "[1,2,3,2,null,2,4]"
	root := BuildTree(input)
	res := removeLeafNodes(root, 2)
	fmt.Println(TreeToArray(res))
}
