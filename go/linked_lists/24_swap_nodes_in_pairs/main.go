// https://leetcode.com/problems/swap-nodes-in-pairs/description/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6}
	head := BuildList(values)
	head = swapPairs(head)
	fmt.Println(LinkedListToArray(head))
}

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
