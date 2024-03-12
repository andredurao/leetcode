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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	nodeStack := []int{}
	totalStack := []int{}

	for head != nil {
		if head.Val == 0 {
			continue
		}
		nodeStack = append(nodeStack, head.Val)
		if len(totalStack) > 0 && ((head.Val > 0 && totalStack[len(totalStack)-1] > 0) || (head.Val < 0 && totalStack[len(totalStack)-1] < 0)) {
			totalStack[len(totalStack)-1] += head.Val
		} else {
			totalStack = append(totalStack, head.Val)
		}
		// fmt.Println("checking:", head.Val, nodeStack, totalStack)
		if len(totalStack) > 1 && (totalStack[len(totalStack)-1]+totalStack[len(totalStack)-2] == 0) {
			// fmt.Println("TOTAL DEL RIGHT")
			lastSum := totalStack[len(totalStack)-1]

			cursor := 0
			total := 0
			for cursor = len(nodeStack) - 1; total != lastSum; cursor-- {
				total += nodeStack[cursor]
			}

			nodeStack = nodeStack[:cursor+1]
			totalStack = totalStack[:len(totalStack)-1]
			// fmt.Println("deleted right val:", head.Val, nodeStack, totalStack)

			// fmt.Println("TOTAL DEL LEFT")
			lastSum = totalStack[len(totalStack)-1]
			total = 0
			for cursor = len(nodeStack) - 1; total != lastSum; cursor-- {
				total += nodeStack[cursor]
			}
			nodeStack = nodeStack[:cursor+1]
			totalStack = totalStack[:len(totalStack)-1]
			// fmt.Println("deleted left val:", head.Val, nodeStack, totalStack)
		}

		// r and l are + and -
		if len(totalStack) > 1 && (nodeStack[len(nodeStack)-1]+nodeStack[len(nodeStack)-2] == 0) {
			// fmt.Println("NODESTACK DEL RIGHT")

			nodeStack = nodeStack[:len(nodeStack)-1]
			totalStack = totalStack[:len(totalStack)-1]
			// fmt.Println("deleted right val:", head.Val, nodeStack, totalStack)

			// fmt.Println("NODESTACK DEL LEFT")
			if totalStack[len(totalStack)-1]+head.Val > 0 {
				totalStack[len(totalStack)-1] += head.Val
			} else {
				totalStack = totalStack[:len(totalStack)-1]
			}

			if nodeStack[len(nodeStack)-1]+head.Val > 0 {
				nodeStack[len(nodeStack)-1] += head.Val
			} else {
				nodeStack = nodeStack[:len(nodeStack)-1]
			}

			// fmt.Println("deleted left val:", head.Val, nodeStack, totalStack)
		}

		head = head.Next
	}

	if len(nodeStack) == 0 {
		return nil
	}

	head = &ListNode{nodeStack[0], nil}
	prev := head
	cur := head

	for i, val := range nodeStack {
		if i > 0 {
			cur = &ListNode{val, nil}
			prev.Next = cur
		}
		prev = cur
	}

	return head
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
