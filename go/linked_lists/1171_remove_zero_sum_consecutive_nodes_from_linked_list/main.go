// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	values := []int{1, 2, -3, 3, 1}
	pos := -1

	head := buildList(values, pos)

	res := removeZeroSumSublists(head)
	printList(res)
}

// ------------------------------------------------------------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	var prev *ListNode
	prev = nil
	stack := []*ListNode{}

	for head != nil {
		fmt.Println("checking:", head.Val)
		// [1, 2, -3, 3, 1]
		//     |
		// check if next two values are present
		if head.Next != nil && (head.Val+head.Next.Val == 0) {
			// fmt.Printf("match %d, %d, %v\n", head.Val, head.Next.Val, stack)
			if len(stack) > 0 {
				prev, stack = stack[len(stack)-1], stack[:len(stack)-1]
			} else {
				prev = nil
			}
			next := head.Next.Next
			head = prev
			if head != nil {
				head.Next = next
			}
		} else {
			stack = append(stack, head)
			head = head.Next
		}
	}

	if len(stack) > 0 {
		return stack[0]
	}
	return nil
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
