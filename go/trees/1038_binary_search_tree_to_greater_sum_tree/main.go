// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/

package main

import (
	"fmt"
	"sort"
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
func bstToGst(root *TreeNode) *TreeNode {
	vals := []int{}
	traverse(root, &vals)
	sum := 0
	for i := range vals {
		sum += vals[i]
	}
	sums := make([]int, len(vals))
	copy(sums, vals)
	for i := range sums {
		sum -= sums[i]
		sums[i] += sum
	}
	// fmt.Println(vals)
	// fmt.Println(sums)
	update(root, &vals, &sums)
	return root
}

func traverse(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, vals)
	*vals = append(*vals, node.Val)
	traverse(node.Right, vals)
}

func update(node *TreeNode, vals *[]int, sums *[]int) {
	if node == nil {
		return
	}
	update(node.Left, vals, sums)
	index := sort.SearchInts(*vals, node.Val)
	node.Val = (*sums)[index]
	update(node.Right, vals, sums)
}

// ------------------------------------------------------

func main() {
	input := "[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]"
	root := BuildTree(input)
	res := bstToGst(root)
	fmt.Println(TreeToArray(res))
}
