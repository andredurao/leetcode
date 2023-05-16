// https://leetcode.com/problems/design-hashmap/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// ["MyHashMap","put","put","get","get","put","get","remove","get"]
	// [[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)

	var1 := obj.Get(1)
	fmt.Println(var1)

	var3 := obj.Get(3)
	fmt.Println(var3)

	obj.Put(2, 1)

	var2 := obj.Get(2)
	fmt.Println(var2)

	obj.Remove(2)

	var2 = obj.Get(2)
	fmt.Println(var2)
}

const Size = 100

type MyHashMap struct {
	Map []*list.List
}

type Node struct {
	key   int
	value int
}

func Constructor() MyHashMap {
	myHashMap := new(MyHashMap)
	myHashMap.Map = make([]*list.List, 100)
	for i := 0; i < Size; i++ {
		myHashMap.Map[i] = list.New()
	}
	return *myHashMap
}

func (this *MyHashMap) Put(key int, value int) {
	page := key % Size
	// Search for element
	for e := this.Map[page].Front(); e != nil; e = e.Next() {
		node, _ := e.Value.(Node)
		if node.key == key {
			// update
			e.Value = Node{key, value}
			return
		}
	}
	this.Map[page].PushBack(Node{key, value})
}

func (this *MyHashMap) Get(key int) int {
	page := key % Size
	// Search for element
	for e := this.Map[page].Front(); e != nil; e = e.Next() {
		node, _ := e.Value.(Node)
		if node.key == key {
			return node.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	page := key % Size
	// Search for element
	for e := this.Map[page].Front(); e != nil; e = e.Next() {
		node, _ := e.Value.(Node)
		if node.key == key {
			this.Map[page].Remove(e)
		}
	}
}

// /**
//  * Your MyHashMap object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Put(key,value);
//  * param_2 := obj.Get(key);
//  * obj.Remove(key);
//  */
