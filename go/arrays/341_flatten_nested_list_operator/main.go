package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }t
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	List  []int
	Index int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	flatList := make([]int, 0)
	buildList(nestedList, &flatList)
	//   fmt.Printf("%v\n", flatList)
	result := &NestedIterator{flatList, 0}
	return result
}

func buildList(listItem []*NestedInteger, flatList *[]int) {
	for _, val := range listItem {
		if val.IsInteger() {
			*flatList = append(*flatList, val.GetInteger())
		} else {
			buildList(val.GetList(), flatList)
		}
	}
}

func (this *NestedIterator) Next() int {
	result := this.List[this.Index]
	this.Index++
	return result
}

func (this *NestedIterator) HasNext() bool {
	return this.Index < len(this.List)
}
