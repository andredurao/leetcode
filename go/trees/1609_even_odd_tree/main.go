// https://leetcode.com/problems/even-odd-tree/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	root := buildTree("[5,4,2,3,3,7]")
	vals, valid := getVals(root)
	fmt.Println(vals, valid)

	// root := buildTree("[5,4,2,3,3,7]")
	// result := isEvenOddTree(root)
	// fmt.Println(result)

}

// -------------------------

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	_, valid := getVals(root)

	return valid
}

func getVals(root *TreeNode) ([][]int, bool) {
	res := [][]int{{}}
	valid := true

	traverse(root, 0, &res, &valid)

	return res, valid
}

func traverse(node *TreeNode, level int, vals *[][]int, valid *bool) {
	if node == nil || !*valid {
		return
	}

	*valid = *valid && (node.Val%2 != level%2)

	// fmt.Println("node =", node.Val, "level =", level, *valid)
	// add a new empty array for the level that we are visiting
	if level+1 > len(*vals) {
		*vals = append(*vals, []int{})
	}

	if level%2 == 0 { // increasing order
		if len((*vals)[level]) > 0 {
			// fmt.Println("inc", node.Val, ">", (*vals)[level][len((*vals)[level])-1])
			*valid = *valid && (node.Val > (*vals)[level][len((*vals)[level])-1])
		}
	} else { // decreasing order
		if len((*vals)[level]) > 0 {
			// fmt.Println("dec", node.Val, "<", (*vals)[level][len((*vals)[level])-1])
			*valid = *valid && (node.Val < (*vals)[level][len((*vals)[level])-1])
		}
	}

	if !*valid {
		return
	}

	(*vals)[level] = append((*vals)[level], node.Val)

	traverse(node.Left, level+1, vals, valid)
	traverse(node.Right, level+1, vals, valid)
}

// -------------------------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(input string) *TreeNode {
	if input == "[]" {
		return nil
	}
	input = strings.Trim(input, "[]")
	values := strings.Split(input, ",")
	nodes := make([]*TreeNode, len(values))
	for i, strValue := range values {
		if strValue == "null" {
			nodes[i] = nil
			continue
		}
		value, _ := strconv.Atoi(strValue)
		nodes[i] = &TreeNode{value, nil, nil}
	}
	i := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}
		i++
		if i >= len(nodes) {
			break
		}
		node.Left = nodes[i]
		i++
		if i >= len(nodes) {
			break
		}
		node.Right = nodes[i]
	}
	return nodes[0]
}
