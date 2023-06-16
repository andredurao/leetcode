// https://leetcode.com/problems/find-mode-in-binary-search-tree/

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
	root := buildTree("[1,null,2,2]")
	result := findMode(root)
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
func findMode(root *TreeNode) []int {
	freqMap := make(map[int]int, 0)
	walk(root, freqMap)
	max := 0
	for _, count := range freqMap {
		if count > max {
			max = count
		}
	}

	result := make([]int, 0)

	for val, count := range freqMap {
		if count == max {
			result = append(result, val)
		}
	}

	return result
}

func walk(node *TreeNode, freqMap map[int]int) {
	if node == nil {
		return
	}
	_, found := freqMap[node.Val]
	if found {
		freqMap[node.Val]++
	} else {
		freqMap[node.Val] = 1
	}
	walk(node.Left, freqMap)
	walk(node.Right, freqMap)
}

// Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).
