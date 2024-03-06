// https://leetcode.com/problems/linked-list-cycle/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	values := []int{1, 2}
	pos := -1

	head := buildList(values, pos)

	fmt.Println(hasCycle(head))
}

// The previous ruby version I could define a new instance variable, but using go that's not possible
// I've used a map of *ListNode to store the visited nodes

// ------------------------------------------------------------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	visitMap := map[*ListNode]struct{}{}

	for head != nil {
		_, found := visitMap[head]
		if !found {
			visitMap[head] = struct{}{}
		} else {
			return true
		}
		head = head.Next
	}
	return false
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
