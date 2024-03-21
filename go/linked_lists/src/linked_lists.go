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
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func LinkedListToArray(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
