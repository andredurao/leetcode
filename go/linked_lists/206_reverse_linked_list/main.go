// https://leetcode.com/problems/reverse-linked-list/description/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	head := BuildList(values)
	res := reverseList(head)
	fmt.Println(LinkedListToArray(res))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	stack := []int{}
	for head != nil {
		stack = append([]int{head.Val}, stack...)
		head = head.Next
	}

	head = &ListNode{stack[0], nil}
	first := head
	for i := 1; i < len(stack); i++ {
		head.Next = &ListNode{stack[i], nil}
		head = head.Next
	}

	return first
}
