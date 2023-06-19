// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/description/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	root := buildTree("[4,8,5,0,1,null,6]")
	result := averageOfSubtree(root)
	fmt.Println(result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
	total := 0
	walk(root, &total)
	return total
}

func walk(node *TreeNode, total *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftTotal, leftCount := walk(node.Left, total)
	rightTotal, rightCount := walk(node.Right, total)
	currTotal := (leftTotal + rightTotal + node.Val)
	currCount := (leftCount + rightCount + 1)
	avg := currTotal / currCount
	// fmt.Println("val", node.Val, "tot", currTotal, "count", currCount, "avg", avg, "gtot", *total)
	if node.Val == avg {
		*total++
	}
	return currTotal, currCount
}
