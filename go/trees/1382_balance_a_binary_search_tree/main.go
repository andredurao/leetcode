// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/

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
func balanceBST(root *TreeNode) *TreeNode {
	vals := []int{}
	traverse(root, &vals)

	return populateTree(&vals, 0, len(vals)-1)
}

func traverse(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, vals)
	*vals = append(*vals, node.Val)
	traverse(node.Right, vals)
}

func populateTree(vals *[]int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	leftNode := populateTree(vals, l, mid-1)
	rightNode := populateTree(vals, mid+1, r)
	return &TreeNode{Val: (*vals)[mid], Left: leftNode, Right: rightNode}
}

// ------------------------------------------------------

func main() {
	input := "[1,null,2,null,3,null,4]"
	root := BuildTree(input)
	res := balanceBST(root)
	fmt.Println(TreeToArray(res))
}
