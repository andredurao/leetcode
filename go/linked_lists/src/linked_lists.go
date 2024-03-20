// https://leetcode.com/problems/linked-list-cycle/description/

package LinkedLists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(values []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	var prev *ListNode

	if len(values) == 0 {
		return head
	}
	head = &ListNode{values[0], nil}
	tail = head

	for i, val := range values {
		if i > 0 {
			tail = &ListNode{val, nil}
			prev.Next = tail
		}
		prev = tail
	}

	return head
}

func PrintList(head *ListNode) {
	i := 0
	for head != nil && i < 10 {
		fmt.Println("val", head.Val)
		head = head.Next
		i++
	}
}
