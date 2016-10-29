package main

// here i'll practice stuff with channels
import (
	"fmt"
	"time"
)

var Done chan bool

// 1. Two goroutines send each other messages
// 2. If at least one is over x, send to Done channel

type Dude struct {
	input chan int
	clock int
	id    int
}

func (d *Dude) send(recv *Dude) {
	d.clock++
	fmt.Println(d.id, "Sending clock: ", d.clock)
	recv.input <- d.clock
	time.Sleep(time.Millisecond * 900)
}

func (d *Dude) receive(sender *Dude) {
	for i := 0; i < 3; i++ {
		select {
		case <-d.input:
			n := <-d.input
			fmt.Println(d.id, ": receiving: ", n)
			d.clock = n
			if d.clock > 3 {
				//Done <- true
				return
			}
			// else, send message back
			//d.send(sender)
		default:
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func nuDude(set int) *Dude {
	ret := new(Dude)
	ret.input = make(chan int)
	ret.clock = 0
	ret.id = set
	return ret
}

func every3(times int, c chan string, done chan bool) {
	//n := 0
	for i := 0; i < times; i++ {
		if i%3 == 0 {
			c <- "Got 3"
		}
	}
	done <- true
}

func catch3(c chan string, done chan bool) {
	n := 0
	for {
		select {
		case msg := <-c:
			n++
			fmt.Println(msg, n)
		case <-done:
			fmt.Println("Finished")
			return
		}
	}
}

func exchange(out chan string) {
	//for i := 0; i < 3; i++ {
	i := 1
	out <- fmt.Sprintf("Hello, %d", i)
	time.Sleep(time.Millisecond * 500)
	/*select {
	case <-in:
		fmt.Println("IN: ", <-in)
	}*/
	//}
}

func exchanger(in chan string) {
	for {
		select {
		case <-in:
			fmt.Println("IN: ", <-in)
		}
	}
}

func main() {
	/*Done = make(chan bool, 2)
	d1 := nuDude(1)
	d2 := nuDude(2)
	go d1.receive(d2)
	go d2.receive(d1)
	d1.send(d2)
	go d2.send(d1)
	go d1.send(d2)*/

	/*select {
	case <-Done:
		fmt.Println("FINISHING")
	}*/
	// ======================
	/*
		c := make(chan string)
		done := make(chan bool)
		go every3(10, c, done)
		catch3(c, done)
	*/
	// ======================

	//in := make(chan string)
	out := make(chan string)
	go exchanger(out)
	exchange(out)

	time.Sleep(time.Millisecond * 1000)
	//fmt.Println("d1 -> ", d1.clock, " d2 -> ", d2.clock)
}
