// https://leetcode.com/problems/diameter-of-binary-tree/description/
// from a previously solution written in ruby

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	vals := [3]int{}
	vals[0] = diameterOfBinaryTree(root.Left)
	vals[1] = diameterOfBinaryTree(root.Right)
	vals[2] = legspan(root)

	res := 0
	for _, val := range vals {
		if val > res {
			res = val
		}
	}

	return res
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := height(node.Left)
	r := height(node.Right)

	if l >= r {
		return l + 1
	} else {
		return r + 1
	}
}

func legspan(node *TreeNode) int {
	return height(node.Left) + height(node.Right)
}

// ------------------------------------------------------

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

func main() {
	root := buildTree("[1,2,3,4,5]")
	fmt.Println(diameterOfBinaryTree(root))
}
