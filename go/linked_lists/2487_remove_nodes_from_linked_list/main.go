//

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	// values := []int{5, 2, 13, 3, 8}
	values := []int{5, 2, 13, 3, 8, 9}
	// values := []int{1, 1, 1, 1, 1}
	head := BuildList(values)
	res := removeNodes(head)
	fmt.Println(LinkedListToArray(res))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	node := head

	stack := []*ListNode{}

	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}

	res := stack[len(stack)-1]
	for i := len(stack) - 1; i >= 0; i-- {
		if i == len(stack)-1 {
			continue
		}
		if stack[i].Val >= res.Val {
			tmp := res
			res = stack[i]
			res.Next = tmp
		}
	}

	return res
}
