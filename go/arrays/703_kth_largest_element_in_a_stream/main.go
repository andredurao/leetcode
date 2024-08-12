//

package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))
}

// -------------------------------------

type KthLargest struct {
	K      int
	Stream []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.IntSlice.Sort(nums)
	return KthLargest{k, nums}
}

func (this *KthLargest) Add(val int) int {
	i := sort.SearchInts(this.Stream, val)
	this.Stream = slices.Insert(this.Stream, i, val)
	return this.Stream[len(this.Stream)-this.K]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
