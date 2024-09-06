// https://leetcode.com/problems/swap-nodes-in-pairs/description/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	// values := []int{1, 2, 3, 4, 5}
	// head := BuildList(values)
	// res := modifiedList([]int{1, 2, 3}, head)

	// values := []int{1, 2, 1, 2, 1, 2}
	// head := BuildList(values)
	// res := modifiedList([]int{1}, head)

	values := []int{1, 2, 3, 4}
	head := BuildList(values)
	res := modifiedList([]int{5}, head)

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
func modifiedList(nums []int, head *ListNode) *ListNode {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}

	var res, cur *ListNode

	for head != nil {
		_, found := set[head.Val]
		if !found {
			// fmt.Println("NF")
			if cur == nil { // 1st node only
				res = head
				cur = res
			} else {
				// fmt.Println("APPEND")
				cur.Next = head
				cur = cur.Next
			}
		}
		head = head.Next
	}
	cur.Next = nil

	return res
}
