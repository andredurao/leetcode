// https://leetcode.com/problems/merge-in-between-linked-lists/

package main

import . "linked_lists"

func main() {
	values1 := []int{10, 1, 13, 6, 9, 5}
	a := 3
	b := 4
	values2 := []int{1000000, 1000001, 1000002}
	list1 := BuildList(values1)
	list2 := BuildList(values2)

	res := mergeInBetween(list1, a, b, list2)
	PrintList(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	first := list1
	i := 0
	aNode := &ListNode{}
	bNode := &ListNode{}
	// [0 1 2 3 4 5 6 7] 3, 4
	for i <= b {
		if i == a-1 {
			aNode = list1
		}
		list1 = list1.Next
		i++
	}
	bNode = list1

	aNode.Next = list2

	for aNode.Next != nil {
		aNode = aNode.Next
	}

	aNode.Next = bNode

	return first
}
