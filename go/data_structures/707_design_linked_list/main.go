// https://leetcode.com/problems/design-linked-list/

package main

import "fmt"

func main() {
	obj := Constructor()
	// fmt.Println(obj)

	// obj.AddAtHead(3)
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)

	// tmp := obj.Get(0)
	// fmt.Println(tmp)

	// // obj.AddAtHead(val);
	// obj.AddAtTail(4)
	// tmp = obj.Get(3)
	// fmt.Println(tmp)

	// tmp = obj.Get(4)
	// fmt.Println(tmp)

	// obj.Debug()

	// obj.AddAtIndex(4, 5)
	// obj.Debug()
	// obj.AddAtIndex(8, 5)
	// obj.Debug()

	// obj.DeleteAtIndex(4)
	// obj.Debug()

	obj.AddAtHead(84)
	obj.AddAtTail(2)
	obj.AddAtTail(39)
	obj.Get(3)
	obj.Get(1)
	obj.AddAtTail(42)
	obj.AddAtIndex(1, 80)
	obj.AddAtHead(14)
	obj.AddAtHead(1)
	obj.AddAtTail(53)
	obj.AddAtTail(98)
	obj.AddAtTail(19)
	obj.AddAtTail(12)

	// 1 14 84 80 2 39 42 53 98 19 12

	obj.Debug()
	fmt.Println(obj.Length)
	// 1 14 84 80 2 39 42 53 98 19 12
	// 10

	obj.Get(2)
	obj.AddAtHead(16)
	obj.AddAtHead(33)
	obj.AddAtIndex(4, 17)
	obj.AddAtIndex(6, 8)
	obj.AddAtHead(37)
	obj.AddAtTail(43)
	obj.DeleteAtIndex(11)
	obj.AddAtHead(80)
	obj.AddAtHead(31)
	obj.AddAtIndex(13, 23)
	obj.AddAtTail(17)
	obj.Get(4)
	obj.AddAtIndex(10, 0)
	obj.AddAtTail(21)
	obj.AddAtHead(73)
	obj.AddAtHead(22)
	obj.AddAtIndex(24, 37)
	obj.AddAtTail(14)
	obj.AddAtHead(97)
	obj.AddAtHead(8)
	obj.Get(6)
	obj.DeleteAtIndex(17)
	obj.AddAtTail(50)
	obj.AddAtTail(28)
	obj.AddAtHead(76)
	obj.AddAtTail(79)
	obj.Get(18)
	obj.DeleteAtIndex(30)
	obj.AddAtTail(5)
	obj.AddAtHead(9)
	obj.AddAtTail(83)
	obj.DeleteAtIndex(3)
	obj.AddAtTail(40)
	obj.DeleteAtIndex(26)
	obj.AddAtIndex(20, 90)
	obj.DeleteAtIndex(30)
	obj.AddAtTail(40)
	obj.AddAtHead(56)
	obj.AddAtIndex(15, 23)
	obj.AddAtHead(51)
	obj.AddAtHead(21)
	obj.Get(26)
	obj.AddAtHead(83)
	obj.Get(30)
	obj.AddAtHead(12)
	obj.DeleteAtIndex(8)
	obj.Get(4)
	obj.AddAtHead(20)
	obj.AddAtTail(45)
	obj.Get(10)
	obj.AddAtHead(56)
	obj.Get(18)
	obj.AddAtTail(33)
	obj.Get(2)
	obj.AddAtTail(70)
	obj.AddAtHead(57)
	obj.AddAtIndex(31, 24)
	obj.AddAtIndex(16, 92)
	obj.AddAtHead(40)
	obj.AddAtHead(23)
	obj.DeleteAtIndex(26)
	obj.Get(1)
	obj.AddAtHead(92)
	obj.AddAtIndex(3, 78)
	obj.AddAtTail(42)
	obj.Get(18)
	obj.AddAtIndex(39, 9)
	obj.Get(13)
	obj.AddAtIndex(33, 17)

	res := obj.Get(51)
	obj.Debug()
	fmt.Println(obj.Length, res)

	obj.AddAtIndex(18, 95)
	obj.AddAtIndex(18, 33)
	obj.AddAtHead(80)
	obj.AddAtHead(21)
	obj.AddAtTail(7)
	obj.AddAtIndex(17, 46)
	obj.Get(33)
	obj.AddAtHead(60)
	obj.AddAtTail(26)
	obj.AddAtTail(4)
	obj.AddAtHead(9)
	obj.Get(45)
	obj.AddAtTail(38)
	obj.AddAtHead(95)
	obj.AddAtTail(78)
	obj.Get(54)
	obj.AddAtIndex(42, 86)

}

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head   *Node
	Length int
}

func Constructor() MyLinkedList {
	return *new(MyLinkedList)
}

func (this *MyLinkedList) Debug() {
	el := this.Head
	for el != nil {
		fmt.Printf("%d ", el.Val)
		el = el.Next
	}
	fmt.Println("")
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Length || index < 0 {
		return -1
	}

	node := this.Head

	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := new(Node)
	node.Val = val
	if this.Head == nil {
		this.Head = node
	} else {
		node.Next = this.Head
		this.Head = node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := new(Node)
	node.Val = val
	if this.Head == nil {
		this.Head = node
	} else {
		el := this.Head
		for el.Next != nil {
			el = el.Next
		}
		el.Next = node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length || index < 0 {
		return
	}
	node := new(Node)
	node.Val = val
	if index == 0 {
		if this.Head == nil {
			this.Head = node
		} else {
			node.Next = this.Head
			this.Head = node
		}
	} else if index == this.Length {
		if this.Head == nil {
			this.Head = node
		} else {
			el := this.Head
			for el.Next != nil {
				el = el.Next
			}
			el.Next = node
		}
	} else {
		el := this.Head
		for i := 0; i < index-1; i++ {
			el = el.Next
		}
		node.Next = el.Next
		el.Next = node
	}
	this.Length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Length || index < 0 {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
	} else {
		el := this.Head
		for i := 0; i < index-1; i++ {
			el = el.Next
		}
		el.Next = el.Next.Next
	}
	this.Length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
