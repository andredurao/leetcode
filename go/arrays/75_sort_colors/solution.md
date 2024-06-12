# Intuition
I was about to implement [insertion sort](https://en.wikipedia.org/wiki/Insertion_sort), but then I thought that colors are a limited set of values, I could store the frequencies of each color in a hashmap and rewrite the values on the input array but sorted, during implementation I've changed the hashmap to a simple array of 3 integers, for each color.

# Approach

1. Initialize an array of integers (`f`), with size 3
2. Iterate nums and increase `f` at each color, after this we would have the frequency of colors
3. Initialize an index var `i` and iterate the values of `f`
3.1. For each value of `f` rewrite `f[i]` using that value and increment `i`

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
func sortColors(nums []int)  {
    f := []int{0,0,0}
    for _, num := range nums {
        f[num]++
    }

    i := 0
    for num := 0 ; num < 3 ; num++ {
        for j := 0 ; j < f[num] ; j++{
            nums[i] = num
            i++
        }
    }
}
```
