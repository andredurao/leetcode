// ported from ruby version that I've submitted 2 years ago

package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	Values map[int]int
	List   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, found := this.Values[val]

	if found {
		return false
	}

	this.Values[val] = len(this.List)
	this.List = append(this.List, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, found := this.Values[val]

	if !found {
		return false
	}

	last := this.List[len(this.List)-1]
	this.List[index] = last
	this.Values[last] = index
	this.List = this.List[:len(this.List)-1]
	delete(this.Values, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.List))
	return this.List[n]
}

func main() {
	rs := Constructor()
	rs.Insert(1)
	rs.Insert(10)
	rs.Insert(20)
	rs.Insert(30)

	for i := 0; i < 100; i++ {
		fmt.Println(i, rs.GetRandom())
	}
}

/**
* Your RandomizedSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Insert(val);
* param_2 := obj.Remove(val);
* param_3 := obj.GetRandom();
 */
