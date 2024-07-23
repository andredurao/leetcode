// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	test([]int{1, 2, 3, 4, 5}, 2)
	test([]int{1}, 1)
	test([]int{1, 2}, 1)
	test([]int{1, 2}, 2)
	test([]int{1, 2, 3}, 1)
	test([]int{1, 2, 3}, 2)
	test([]int{1, 2, 3}, 3)
}

func test(vals []int, n int) {
	fmt.Println("- - - - - - - -", vals, n)
	head := BuildList(vals)
	res := removeNthFromEnd(head, n)
	fmt.Println("_______________")
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	fmt.Println("")
}

// -------------------------------------------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	zero := &ListNode{Val: 0, Next: head}

	l, r := zero, zero

	for i := 0; i <= n; i++ {
		r = r.Next
	}

	for r != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next

	return zero.Next
}
