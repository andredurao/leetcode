// https://leetcode.com/problems/swap-nodes-in-pairs/description/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	head := BuildList(values)
	reorderList(head)
	fmt.Println(LinkedListToArray(head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	list := []*ListNode{}
	cur := head
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	// rewind cur
	cur = head
	// reorder
	var next *ListNode
	i := len(list) - 1
	for i >= len(list)/2 {
		next = cur.Next

		if cur != list[i] {
			cur.Next = list[i]
			list[i].Next = next
		} else {
			// in case list has odd qty
			cur.Next = nil
			break
		}

		cur = next
		i--
	}
	cur.Next = nil
}
