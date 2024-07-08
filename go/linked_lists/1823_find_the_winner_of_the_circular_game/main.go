//

package main

import "fmt"

func main() {
	res := findTheWinner(3, 1)
	fmt.Println(res)
}

// ----------------------------------------------------

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func findTheWinner(n int, k int) int {
	start := &Node{Val: 1, Prev: nil, Next: nil}
	prev := start
	for i := 1; i < n; i++ {
		node := &Node{Val: i + 1, Prev: prev, Next: nil}
		prev.Next = node
		prev = node
	}
	// last next points to start
	prev.Next = start
	// first prev points to end
	start.Prev = prev

	node := start
	for n > 1 {
		for i := 0; i < k-1; i++ {
			node = node.Next
		}
		tmp := node.Next
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
		node = tmp
		n--
	}
	return node.Val
}
