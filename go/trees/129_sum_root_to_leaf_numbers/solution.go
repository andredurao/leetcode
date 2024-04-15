# Intuition
My first thoughts is that I've had to traverse the tree recursively while storing the path, and when it reaches a leaf node I would calculate the total and store in another array to get the total in the end.
I realized I was storing more data than I needed.

# Approach
1. Initialize an int var (`total`) with zero
2. Using an auxiliary recursive function `traverse` do the following
2.1. If the current node is a leaf sum it's path value to the total value passed by reference, otherwise call the function recursively with its left and right nodes and proper path values

# Complexity
- Time complexity: $$O(n*log(n))$$

- Space complexity: $$O(h)$$, given that h is the tree's height

# Code
```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	total := 0
	traverse(root, root.Val, &total)
	return total
}

func traverse(node *TreeNode, pathTotal int, total *int) {
	if node.Left == nil && node.Right == nil {
		*total += pathTotal
		return
	}
	if node.Left != nil {
		traverse(node.Left, (pathTotal*10 + node.Left.Val), total)
	}
	if node.Right != nil {
		traverse(node.Right, (pathTotal*10 + node.Right.Val), total)
	}
}
```
