# Intuition
Initially I've tried to solve the problem by looking up all greater values for each node, but then I realized that I would need to backtrack and my solution was not clear. I've decided to store the values and use an auxiliar array with the sums already calculated then traverse the tree again to update the values.

# Approach
1. Traverse the tree and append the values of each node into an array (`vals`)
2. Store the `sum` of vals
3. Create another array `sums` with the same values as `vals` and iterate its values while:
3.1. decreasing `sum` for each value
3.2. increment each value with the new value of `sum`
4. Traverse the tree again and on each node perform a binary search for their `index` in the `vals` array, update the node value with the value in `sums` using the same `index`

# Complexity
- Time complexity: $$O(n*Log(n))$$
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: $$O(n)$$
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

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
func bstToGst(root *TreeNode) *TreeNode {
	vals := []int{}
	traverse(root, &vals)
	sum := 0
	for i := range vals {
		sum += vals[i]
	}
	sums := make([]int, len(vals))
	copy(sums, vals)
	for i := range sums {
		sum -= sums[i]
		sums[i] += sum
	}
	// fmt.Println(vals)
	// fmt.Println(sums)
	update(root, &vals, &sums)
	return root
}

func traverse(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, vals)
	*vals = append(*vals, node.Val)
	traverse(node.Right, vals)
}

func update(node *TreeNode, vals *[]int, sums *[]int) {
	if node == nil {
		return
	}
	update(node.Left, vals, sums)
	index := sort.SearchInts(*vals, node.Val)
	node.Val = (*sums)[index]
	update(node.Right, vals, sums)
}
```
