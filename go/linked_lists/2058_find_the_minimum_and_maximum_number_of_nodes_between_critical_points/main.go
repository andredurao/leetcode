// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points

package main

import (
	"fmt"
	. "linked_lists"
)

func main() {
	values := []int{5, 3, 1, 2, 5, 1, 2}
	head := BuildList(values)
	res := nodesBetweenCriticalPoints(head)
	fmt.Println(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	window := []int{-1, -1, -1}
	index, minDist, firstIndex, lastIndex := 0, 999999, -1, -1

	for head != nil {
		moveWindow(&window, head.Val)
		if isCritical(&window) {
			if firstIndex == -1 {
				firstIndex = index
			} else {
				minDist = min(minDist, (index - lastIndex))
			}
			lastIndex = index
		}
		head = head.Next
		index++
	}

	res := []int{-1, -1}

	if minDist < 999999 {
		res[0] = minDist
		res[1] = lastIndex - firstIndex
	}

	return res
}

func isCritical(window *[]int) bool {
	if !validWindow(window) {
		return false
	}
	return ((*window)[1] > (*window)[0] && (*window)[1] > (*window)[2]) ||
		((*window)[1] < (*window)[0] && (*window)[1] < (*window)[2])
}

func validWindow(window *[]int) bool {
	return (*window)[0] != -1 && (*window)[1] != -1 && (*window)[2] != -1
}

func moveWindow(window *[]int, val int) {
	(*window)[0], (*window)[1], (*window)[2] = (*window)[1], (*window)[2], val
}
