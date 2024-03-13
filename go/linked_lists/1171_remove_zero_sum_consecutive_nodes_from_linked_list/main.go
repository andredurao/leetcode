// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	values := []int{0} // 5,-3,-4,1,6,-2,-5
	pos := -1

	head := buildList(values, pos)

	res := removeZeroSumSublists(head)
	printList(res)
}

// ------------------------------------------------------------------

// IGNORE results : based on editorial
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	front := &ListNode{0, head}
	node := front
	total := 0
	totalRefMap := map[int]*ListNode{}

	for node != nil {
		total += node.Val
		lastOccurrence, found := totalRefMap[total]
		if found {
			node = lastOccurrence.Next
			// cleanup previous node refs in hash
			p := total + node.Val
			for p != total {
				delete(totalRefMap, p)
				node = node.Next
				p += node.Val
			}

			// "jump" from lastOccurrence to the next of total
			lastOccurrence.Next = node.Next
		} else {
			totalRefMap[total] = node
		}
		node = node.Next
	}

	return front.Next
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
