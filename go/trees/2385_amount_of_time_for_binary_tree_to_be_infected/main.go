package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "[1,5,3,null,4,10,6,9,2]"
	root := buildTree(str)

	res := amountOfTime(root, 3)
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
func amountOfTime(root *TreeNode, start int) int {
	adjacencies := map[int]map[int]struct{}{}

	visit(root, nil, adjacencies)

	return bfs(root, start, 9, adjacencies)
}

func bfs(root *TreeNode, start int, end int, adjacencies map[int]map[int]struct{}) int {
	d := 0
	q := []int{}
	explored := map[int]int{}
	explored[start] = 0
	q = append(q, start)
	for len(q) > 0 {
		// dequeue
		v := q[len(q)-1]
		q = q[:len(q)-1]
		for w := range adjacencies[v] {
			d++
			_, found := explored[w]
			if !found {
				explored[w] = explored[v] + 1
				q = append(q, w)
			}
		}
	}
	// lookup for largest distance
	max := 0
	for _, d := range explored {
		if d > max {
			max = d
		}
	}
	return max
}

// Visit nodes filling up adjacency map
func visit(node *TreeNode, parent *TreeNode, adjacencies map[int]map[int]struct{}) {
	if node == nil {
		return
	}
	_, found := adjacencies[node.Val]
	if !found {
		adjacencies[node.Val] = map[int]struct{}{}
	}
	if parent != nil {
		adjacencies[parent.Val][node.Val] = struct{}{}
		adjacencies[node.Val][parent.Val] = struct{}{}
	}
	visit(node.Left, node, adjacencies)
	visit(node.Right, node, adjacencies)
}
