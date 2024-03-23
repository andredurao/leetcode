# Intuition
Initially I thought the problem was easy to solve, but later I've realized that it was no so simple as creating a list and building the linked list all over again. That would be in the same lines as changing the values of nodes instead of changing their next references. I only realized later that it was necessary to also have a `prev` var to store the previous current

# Approach
1. Check if head is null or head.next is null, if so there's nothing to swap and just return head as it is
2. Store the following references:
2.1. `prev` as null
2.2. `next` as head.next
2.3. `res` (the result) equals to `next`
3. Repeat while `head` and `next` are not null:
3.1. Set `prev.next` as `next` only if `prev` is present
3.2. Swap: `head.next` points to `next.next` and `next.next` points to `head`
3.3. Move pointers:
    - Set `prev` as `head`
    - Set `head` as `head.next`
    - Set `next` as `head.next`, which has just been updated, only if head is present. Otherwise set it as null
4. Return `res`

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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, res, next *ListNode
	next = head.Next
	res = next

	for head != nil && next != nil {
		// swap
		if prev != nil {
			prev.Next = next
		}
		head.Next = next.Next
		next.Next = head

		// move pointers
		prev = head
		head = head.Next
		if head != nil {
			next = head.Next
		} else {
			head = nil
		}
	}

	return res
}

```
