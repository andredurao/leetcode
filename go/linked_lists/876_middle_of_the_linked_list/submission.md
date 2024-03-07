# Intuition
To know the middle of a linked list, I need to get the length of the list

# Approach
1. Store head in a different pointer (node)
2. Traverse the list incrementing an int var using node until node is null
3. Move head to it's next Node length / 2 times

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	l := 0
	var node *ListNode
	node = head
	for node != nil {
		node = node.Next
		l++
	}

	for i := 0; i < (l / 2); i++ {
		head = head.Next
	}
	return head
}

```
