// https://leetcode.com/problems/reverse-linked-list/description/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	values := []int{1, 2, 2, 1}
	head := BuildList(values)
	res := isPalindrome(head)
	fmt.Println(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	values := []int{}

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-i-1] {
			return false
		}
	}

	return true
}
