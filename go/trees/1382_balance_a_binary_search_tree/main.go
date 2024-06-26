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

	l, r := 0, len(vals)-1
	midIndex := (l + r) / 2
	mid := &TreeNode{Val: vals[midIndex], Left: nil, Right: nil}
	used := map[int]struct{}{}
	used[mid.Val] = struct{}{}
	populateTree(mid, &vals, l, r, used)

	return mid
}

func traverse(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, vals)
	*vals = append(*vals, node.Val)
	traverse(node.Right, vals)
}

func populateTree(node *TreeNode, vals *[]int, l, r int, used map[int]struct{}) *TreeNode {
	if l < 0 || r >= len(*vals) {
		return nil
	}
	mid := (l + r) / 2
	if node == nil {
		if _, found := used[(*vals)[mid]]; found {
			return nil
		} else {
			node = &TreeNode{Val: (*vals)[mid]}
			used[node.Val] = struct{}{}
		}
	}
	node.Left = populateTree(node.Left, vals, l, mid, used)
	node.Right = populateTree(node.Right, vals, mid+1, r, used)
	return node
}

// ------------------------------------------------------

func main() {
	input := "[1,null,2,null,3,null,4]"
	root := BuildTree(input)
	res := balanceBST(root)
	fmt.Println(TreeToArray(res))
}
