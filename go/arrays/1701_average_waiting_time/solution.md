# Intuition
My first thought was that I've had to simulate the customers queue and accumulate the sum of waiting times to get the average at the end.

# Approach
1. Initialize to int variables: time and totalTime with zero
2. Iterate through the customers array and update time to the customer arrival if that was after the current time.
2.1. Get the finish time of the order, by adding the second item of the order to the current time and update time to that value.
2.2. Accumulate the value $$finish - customer[0]$$ into totalTime
3. Return the average of $$totalTime$$ by the qty of customers

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
func averageWaitingTime(customers [][]int) float64 {
    totalTime := 0
    time := 0
	for _, customer := range customers {
        if customer[0] > time {
		    time = customer[0]
        }
		finish := time + customer[1]
        time = finish
        totalTime += finish - customer[0]
	}

	return float64(totalTime) / float64(len(customers))
}
```
