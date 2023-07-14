package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "[10,5,15,3,7,null,18]"
	root := buildTree(str)
	res := rangeSumBST(root, 7, 15)
	fmt.Println(res)
}

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	traverse(root, &low, &high, &sum)
	return sum
}

func traverse(node *TreeNode, low *int, high *int, sum *int) {
	if node == nil {
		return
	}
	traverse(node.Left, low, high, sum)
	traverse(node.Right, low, high, sum)

	if node.Val >= *low && node.Val <= *high {
		*sum += node.Val
	}
}
