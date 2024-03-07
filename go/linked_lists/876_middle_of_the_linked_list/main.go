// https://leetcode.com/problems/linked-list-cycle/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6}
	pos := -1

	head := buildList(values, pos)

	mid := middleNode(head)
	printList(mid)

}

// ------------------------------------------------------------------

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

// ------------------------------------------------------------------

func buildList(values []int, pos int) *ListNode {
	var tailRef *ListNode
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
		if i == pos {
			tailRef = tail
		}
		prev = tail
	}

	if tailRef != nil {
		tail.Next = tailRef
	}

	return head
}

func printList(head *ListNode) {
	i := 0
	for head != nil && i < 10 {
		fmt.Println("val", head.Val)
		head = head.Next
		i++
	}
}
