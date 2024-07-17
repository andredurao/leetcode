// https://leetcode.com/problems/delete-nodes-and-return-forest/

package main

import (
	"fmt"
	. "tree_node"
)

func main() {
	// input := "[1,2,3,4,5,6,7]"
	input := "[1,2,3,null,null,null,4]"

	root := BuildTree(input)
	to_delete := []int{2, 1}
	res := delNodes(root, to_delete)
	for _, t := range res {
		fmt.Println(TreeToArray(t))
	}
}

// --------------------------------------------------------------------
type DelNodeStruct struct {
	Parent          *TreeNode
	Node            *TreeNode
	MarkedForDelete bool
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	nodesMap := map[int]DelNodeStruct{}
	// populate the map of nodes to know their parents and depth
	traverse(nil, root, nodesMap)
	// mark nodes for removal
	for _, val := range to_delete {
		newDelNodeStruct := nodesMap[val]
		newDelNodeStruct.MarkedForDelete = true
		nodesMap[val] = newDelNodeStruct
	}

	// fmt.Println(nodesMap)

	// delete nodes
	for _, nodeStruct := range nodesMap {
		parent := nodeStruct.Parent
		if nodeStruct.MarkedForDelete {
			// fmt.Println("del", nodeStruct.Node.Val)
			if parent != nil {
				if parent.Left == nodeStruct.Node {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}
			if nodeStruct.Node.Left != nil {
				tmp := nodesMap[nodeStruct.Node.Left.Val]
				tmp.Parent = nil
				nodesMap[nodeStruct.Node.Left.Val] = tmp
			}
			if nodeStruct.Node.Right != nil {
				tmp := nodesMap[nodeStruct.Node.Right.Val]
				tmp.Parent = nil
				nodesMap[nodeStruct.Node.Right.Val] = tmp
			}
		}
	}

	// Add items that were not deleter and doesn't have a parent
	res := []*TreeNode{}
	for _, nodeStruct := range nodesMap {
		parent := nodeStruct.Parent
		if parent == nil && !nodeStruct.MarkedForDelete {
			// fmt.Println("add", nodeStruct.Node.Val)
			res = append(res, nodeStruct.Node)
		}
	}

	return res
}

func traverse(parent, node *TreeNode, nodesMap map[int]DelNodeStruct) {
	if node == nil {
		return
	}
	nodesMap[node.Val] = DelNodeStruct{Parent: parent, Node: node, MarkedForDelete: false}
	traverse(node, node.Left, nodesMap)
	traverse(node, node.Right, nodesMap)
}
