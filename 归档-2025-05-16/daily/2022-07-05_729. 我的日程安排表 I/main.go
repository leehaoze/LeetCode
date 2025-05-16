package main

import (
	"fmt"
)

type Event struct {
	start int
	end   int
	next  *Event
}

type MyCalendar struct {
	head *Event
}

func Constructor() MyCalendar {
	return MyCalendar{head: &Event{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if !this.check(start, end) {
		return false
	}

	this.set(start, end)
	return true
}

func (this *MyCalendar) set(start, end int) {
	event := &Event{
		start: start,
		end:   end,
		next:  nil,
	}

	// 寻找插入点  nodeA--newEvent--nodeB
	// 要求 newEvent.start >= nodeA.end && newEvent.end <= nodeB.start
	point := this.head
	for point.next != nil {
		point = point.next
	}

	point.next = event
}

func (this *MyCalendar) check(start, end int) bool {
	event := this.head.next
	for event != nil {
		if (start >= event.start && start < event.end) || (end > event.start && end < event.end) || (start <= event.start && end >= event.end) {
			return false
		}

		event = event.next
	}

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20))
	fmt.Println(obj.Book(15, 25))
	fmt.Println(obj.Book(20, 30))
}
