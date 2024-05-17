# Intuition

I've checked the description hint about rebuilding the tree and while I was implementing it I've decided to try to DFS the tree while removing the previous to current reference when a node matches the conditions to be deleted (leaf and with target value), but since there's no "parent" reference in this structure I thought that I could traverse and pass the "parent" as an argument all the way to a leaf node recursively and pop back the recursive tree while deleting references when possible. The only case which this idea would not work was the minimum case when root itself has no "children" and it's value is the target, for that I've decided to return null in the original function of the problem.

# Approach
1. Save the original root reference to a temporary value
2. Call a recursive function (fn) with the args: previous, current, direction and target, as (null, root, -1, target) that does the following:
2.1. Call fn recursively until the current node is null with the current node and it's left and right references identified by a `dir` argument, 0 for left and 1 for right
2.2. If the current node matches the delete conditions, set the previous reference to it using the `dir` argument to null
3. Return null if only one node remains and that also matches the conditions, otherwise return the temporary value from [1]

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    tmp := root
    traverseAndDelete(nil, root, -1, &target)

    // single node with target value case
    if tmp != nil && tmp.Left == nil && tmp.Right == nil && tmp.Val == target {
        return nil
    }
    return tmp
}

func traverseAndDelete(prev *TreeNode, cur *TreeNode, dir int, target *int) {
    if cur == nil {
        return
    }

    traverseAndDelete(cur, cur.Left, 0, target)
    traverseAndDelete(cur, cur.Right, 1, target)

    if prev != nil && cur.Left == nil && cur.Right == nil && cur.Val == *target {
        if dir == 0 { // Left
            prev.Left = nil
        } else if dir == 1 {
            prev.Right = nil
        }
        return
    }
}
```
