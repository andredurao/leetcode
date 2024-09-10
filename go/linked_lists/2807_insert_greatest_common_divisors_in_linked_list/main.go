// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	// values := []int{18, 6, 10, 3}
	values := []int{7}
	head := BuildList(values)
	res := insertGreatestCommonDivisors(head)

	fmt.Println(LinkedListToArray(res))
}

// ---------------------------------------------------------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	res := head

	for head != nil {
		if head.Next != nil {
			fmt.Println("insert GCD", head.Val, head.Next.Val, gcd(head.Val, head.Next.Val))
			node := &ListNode{Val: gcd(head.Val, head.Next.Val), Next: head.Next}
			head.Next = node
			head = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return res
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
