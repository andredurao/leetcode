// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	test([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	test([][]int{})
	test([][]int{{1}, {0}})
}

func test(input [][]int) {
	lists := []*ListNode{}
	for _, vals := range input {
		lists = append(lists, BuildList(vals))
	}
	res := mergeKLists(lists)
	fmt.Println("_______________")
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	fmt.Println("")
}

// -------------------------------------------------

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{Val: 999999, Next: nil}
	for _, lh := range lists {
		for lh != nil {
			// change 1st element
			if head.Val == 999999 {
				head = lh
				lh = lh.Next
				head.Next = nil
				continue
			}
			node := findInsertNode(head, lh)

			tmp := lh
			lh = lh.Next

			if tmp.Val < head.Val {
				tmp.Next = head
				head = tmp
			} else {
				tmp2 := node.Next
				node.Next = tmp
				tmp.Next = tmp2
			}
		}
	}

	if head.Val == 999999 {
		return nil
	}
	return head
}

func findInsertNode(head, node *ListNode) *ListNode {
	for head.Next != nil && head.Next.Val < node.Val {
		head = head.Next
	}
	return head
}

func printlist(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
