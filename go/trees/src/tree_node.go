// https://leetcode.com/problems/linked-list-cycle/description/

package TreeNode

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

func BuildTree(input string) *TreeNode {
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

func TreeToArray(root *TreeNode) string {
	nodeMap := map[int]string{}
	maxIndex := 0
	traverseFillNodeMap(root, 0, nodeMap, &maxIndex)

	vals := []string{}
	for i := 0; i <= maxIndex; i++ {
		val, found := nodeMap[i]
		if found {
			vals = append(vals, val)
		} else {
			vals = append(vals, "null")
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(vals, ","))
}

func traverseFillNodeMap(node *TreeNode, index int, nodeMap map[int]string, maxIndex *int) {
	if node == nil {
		return
	}
	nodeMap[index] = fmt.Sprintf("%d", node.Val)
	*maxIndex = max(index, *maxIndex)

	traverseFillNodeMap(node.Left, 2*index+1, nodeMap, maxIndex)
	traverseFillNodeMap(node.Right, 2*index+2, nodeMap, maxIndex)
}
