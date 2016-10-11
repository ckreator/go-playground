package main

import "fmt"

// queue functionality:
// add
//

type all interface{}

type Node struct {
	val  all
	next *Node
	prev *Node
}

type Queue struct {
	head *Node
	tail *Node
	len  int
}

func (q *Queue) add(v all) {
	n := Node{val: v}
	fmt.Println("pushing: ", n)
	if q.tail == nil {
		q.tail = &n
		q.head = &n
		return
	}
	q.tail.prev = &n
	q.tail = &n
}

func (q *Queue) addFirst(v all) {
	n := Node{val: v}
	fmt.Println("putting at head: ", n)
	if q.tail == nil {
		q.tail = &n
		q.head = &n
		return
	}
	n.prev = q.head
	q.head.next = &n
	q.head = q.head.next
}

func (q *Queue) get() all {
	ret := q.head
	fmt.Println("polling: ", ret)
	if ret == nil {
		return nil
	}
	q.head = q.head.prev
	if q.head != nil && q.head.next != nil {
		q.head.next = nil
	}
	return ret.val
}

func (q *Queue) getLast() all {
	ret := q.tail
	fmt.Println("polling: ", ret.val)
	if ret == nil {
		return nil
	}
	if q.tail != nil && q.tail.next != nil {
		q.tail = nil
		return ret.val
	}
	//q.tail.next.prev = q.tail.next
	q.tail = q.tail.next
	//q.tail.prev = nil
	return ret.val
}

// stackify will simply revert the queue
func (q *Queue) stackify() {
	if q.head == nil || q.tail == nil {
		return
	}
	//	var tmp *Node = q.head
}

func main() {
	_q := Queue{}
	q := &_q
	q.add(5)
	q.add("a string")
	q.addFirst("should be first")
	q.get()
	q.getLast()
	q.get()
}
