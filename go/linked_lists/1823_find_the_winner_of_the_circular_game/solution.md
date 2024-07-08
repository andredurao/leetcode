# Intuition
Initially I've tried to approach the problem using an array of indexes and use remainder of divisions to find the circular shifts on each iteration, but it didn't work as expected. I decided to implement a double linked list and connect their ends to make it circular.

# Approach
1. Initialize a circular double linked list containing the values from 1 to N, and start at the first node (value = 1)
2. While $$n$$ is greater than one do the following:
2.1. move the cursor to the next node
2.2. remove the reference of the cursor by making the neighbours point to each other
2.3. move the cursor to the next node (saved before the removal)
2.4. decrease $$n$$
3. Return the value in the cursor

# Complexity
- Time complexity: $$O(n*k)$$

- Space complexity: $$O(n)$$

# Code
```
type Node struct {
    Val int
    Prev *Node
    Next *Node
}

func findTheWinner(n int, k int) int {
    start := &Node{Val: 1, Prev: nil, Next: nil}
    prev := start
	for i := 1; i < n; i++ {
        node := &Node{Val: i+1, Prev: prev, Next: nil}
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
```
