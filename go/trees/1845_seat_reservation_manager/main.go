// https://leetcode.com/problems/seat-reservation-manager/description/
// 1st time using GODS package: https://github.com/emirpasic/gods#priorityqueue

package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func main() {
	seatManager := Constructor(5)
	seatManager.Reserve()
	seatManager.Reserve()
	seatManager.Unreserve(2)
	seatManager.Reserve()
	seatManager.Reserve()
	seatManager.Reserve()
	seatManager.Reserve()
	seatManager.Unreserve(5)
}

type SeatManager struct {
	Queue *priorityqueue.Queue
}

func Constructor(n int) SeatManager {
	seatManager := new(SeatManager)
	queue := priorityqueue.NewWith(utils.IntComparator)
	seatManager.Queue = queue

	for i := 0; i < n; i++ {
		seatManager.Queue.Enqueue(i + 1)
	}

	return *seatManager
}

func (this *SeatManager) Reserve() int {
	seat, _ := this.Queue.Dequeue()
	return seat.(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.Queue.Enqueue(seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
