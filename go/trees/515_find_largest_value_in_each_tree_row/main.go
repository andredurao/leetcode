// https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/

package main

import (
	"fmt"
	"math"
)

// TODO: check reader implementation in ruby
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Left.Left = &TreeNode{5, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}

	result := largestValues(root)
	fmt.Printf("%v\n", result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	valuesMap := make(map[int]int, 0)
	maxHeight := 0
	traverse(root, 0, valuesMap, &maxHeight)
	result := make([]int, maxHeight+1)
	for i := 0; i <= maxHeight; i++ {
		result[i], _ = valuesMap[i]
	}

	return result
}

func traverse(node *TreeNode, height int, valuesMap map[int]int, maxHeight *int) {
	if node == nil {
		return
	}
	*maxHeight = int(math.Max(float64(height), float64(*maxHeight)))
	cur, found := valuesMap[height]
	if found {
		max := int(math.Max(float64(cur), float64(node.Val)))
		valuesMap[height] = max
	} else {
		valuesMap[height] = node.Val
	}
	traverse(node.Left, height+1, valuesMap, maxHeight)
	traverse(node.Right, height+1, valuesMap, maxHeight)
}
