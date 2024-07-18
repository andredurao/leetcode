// https://leetcode.com/problems/delete-nodes-and-return-forest/

package main

import (
	"fmt"
	. "tree_node"
)

func main() {
	// input := "[1,2,3,null,4]"
	// input := "[7,1,4,6,null,5,3,null,null,null,null,null,2]"
	// input := "[1,1,1]"
	input := "[1,2,3,4,5,6,7]"

	root := BuildTree(input)
	res := countPairs(root, 3)
	fmt.Println(res)
}

// --------------------------------------------------------------------
type CountPairsStruct struct {
	Parent *TreeNode
	Node   *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
	// populate the map of nodes to know their parents and depth
	nodesMap := map[*TreeNode]*CountPairsStruct{}
	getParentReferences(nil, root, nodesMap, 0)

	counter := 0
	for val, ref := range nodesMap {
		fmt.Println("CUR", val)
		if isLeaf(ref.Node) {
			fmt.Println("LEAF")
			visitMap := map[*TreeNode]struct{}{}
			traverse(ref.Node, ref.Parent, nodesMap, visitMap, 1, distance, &counter)
		}
	}
	return counter / 2
}

func traverse(start *TreeNode, node *TreeNode, nodesMap map[*TreeNode]*CountPairsStruct, visitMap map[*TreeNode]struct{}, distance int, maxDistance int, counter *int) {
	if node == nil || node == start || distance > maxDistance {
		return
	}
	if _, visited := visitMap[node]; visited {
		return
	}
	visitMap[node] = struct{}{}
	if isLeaf(node) {
		fmt.Println("LEAF", "start", start.Val, "end", node.Val, "dist", distance)
		*counter++
		return
	}
	if !isRoot(node, nodesMap) {
		traverse(start, (*nodesMap[node]).Parent, nodesMap, visitMap, distance+1, maxDistance, counter)
	}
	traverse(start, node.Left, nodesMap, visitMap, distance+1, maxDistance, counter)
	traverse(start, node.Right, nodesMap, visitMap, distance+1, maxDistance, counter)
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func isRoot(node *TreeNode, nodesMap map[*TreeNode]*CountPairsStruct) bool {
	return nodesMap[node].Parent == nil
}

func getParentReferences(parent, node *TreeNode, nodesMap map[*TreeNode]*CountPairsStruct, depth int) {
	if node == nil {
		return
	}
	nodesMap[node] = &CountPairsStruct{Parent: parent, Node: node}
	getParentReferences(node, node.Left, nodesMap, depth+1)
	getParentReferences(node, node.Right, nodesMap, depth+1)
}
