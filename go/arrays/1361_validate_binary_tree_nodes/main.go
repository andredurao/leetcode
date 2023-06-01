// https://leetcode.com/problems/validate-binary-tree-nodes/

package main

import "fmt"

func main() {
	n := 4
	leftChild := []int{1, -1, 3, -1}
	rightChild := []int{2, -1, -1, -1}
	result := validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Println(result)

	n = 4
	leftChild = []int{1, -1, 3, -1}
	rightChild = []int{2, 3, -1, -1}
	result = validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Println(result)

	n = 2
	leftChild = []int{1, 0}
	rightChild = []int{-1, -1}
	result = validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Println(result)

	n = 6
	leftChild = []int{1, -1, -1, 4, -1, -1}
	rightChild = []int{2, -1, -1, 5, -1, -1}
	result = validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Println(result)

	n = 4
	leftChild = []int{3, -1, 1, -1}
	rightChild = []int{-1, -1, 0, -1}
	result = validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Println(result)

	n = 4
	leftChild = []int{1, 0, 3, -1}
	rightChild = []int{-1, -1, -1, -1}
	result = validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Println(result)
}

// From editorial
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := findRoot(n, leftChild, rightChild)
	if root == -1 {
		return false
	}

	seen := make(map[int]int, 0)
	stack := make([]int, 0)
	seen[root] = 0
	stack = append(stack, root)

	for len(stack) > 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		children := []int{leftChild[node], rightChild[node]}
		for _, child := range children {
			if child != -1 {
				_, ok := seen[child]
				if ok {
					return false
				}
				stack = append(stack, child)
				seen[child] = 1
			}
		}
	}

	return len(seen) == n
}

func findRoot(n int, leftChild []int, rightChild []int) int {
	children := make(map[int]int)
	for _, l := range leftChild {
		children[l] = 1
	}
	for _, r := range rightChild {
		children[r] = 1
	}

	for i := 0; i < n; i++ {
		_, ok := children[i]
		if !ok {
			return i
		}
	}
	return -1
}
