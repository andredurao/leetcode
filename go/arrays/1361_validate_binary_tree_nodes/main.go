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

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	srcRefArray := make([]int, n)
	for i := 0; i < n; i++ {
		srcRefArray[i] = -1
	}

	for i, l := range leftChild {
		r := rightChild[i]
		// fmt.Println(i, l, r)

		// fmt.Println("l")
		if l != -1 {
			if srcRefArray[l] != -1 {
				return false
			} else {
				srcRefArray[l] = i
			}
		}

		// fmt.Println("r")
		if r != -1 {
			if srcRefArray[r] != -1 {
				return false
			} else {
				srcRefArray[r] = i
			}
		}
	}

	// fmt.Printf("%v\n", srcRefArray)

	rootCount := 0
	for i := 0; i < n; i++ {
		if srcRefArray[i] == -1 {
			rootCount++
		}
		if rootCount > 1 {
			return false
		}
	}
	if rootCount != 1 {
		return false
	}

	//search for cycles
	for i := 0; i < n; i++ {
		path := []int{}
		item := srcRefArray[i]
		for {
			if item == -1 {
				break
			}
			if included(path, item) {
				// fmt.Println("loop")
				return false
			}
			path = append(path, item)
			item = srcRefArray[item]
			// fmt.Printf("i: %d path: %v\n", i, path)
		}
	}

	return true
}

func included(path []int, val int) bool {
	for _, n := range path {
		if val == n {
			return true
		}
	}
	return false
}
