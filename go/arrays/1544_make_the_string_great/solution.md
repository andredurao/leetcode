# Intuition
Initially I thought that I've had to remove all contiguous characters that were the same, like aa, aA, Aa or AA. It took me a while to realize that that only sequences like aA or Aa should be removed. For this I've used a stack and enqueued chars as long as they don't match the condition for removal, otherwise pop the queue

# Approach
1. Initialize an empty queue
2. Iterate string `s` and check if the queue has items and peek the queue to check for the removal conditions (aA or Aa), if so pop the queue, otherwise enqueue the current char

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```
func makeGood(s string) string {
	queue := []rune{}
	for _, ch := range s {
		// for a single toUpper with no conditions: ((ch-97+128)%32 + 65) <- In go the `+128` is needed because go returns negative remainders for `%`
		// but in this case I'm interested in only aA or Aa conditions
		if len(queue) > 0 && ((queue[len(queue)-1]-ch == 32) || (ch-queue[len(queue)-1] == 32)) {
			queue = queue[:len(queue)-1]
		} else {
			queue = append(queue, ch)
		}
	}
	return string(queue)
}

```
