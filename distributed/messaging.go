package main

import (
	"fmt"
	"time"
)

const max_buf = 10

type Wrapped struct {
	timestamp int
	value     interface{}
}

type Finisher struct {
	c         chan bool
	triggered bool
	done      func()
}

var ID = 0
var Done chan bool

type Process struct {
	id     int
	input  chan *Wrapped
	output chan *Wrapped
	done   chan bool
	clock  int
}

func (p *Process) receiver() {
	// get the timestamp
	tm := <-p.input
	fmt.Println("ID: ", p.id, "GOT MESSAGE: ", tm.value)
}

func (p *Process) sender(msg string) {
	tm := new(Wrapped)
	// send the current clock + 1, because process also takes some time
	tm.timestamp = p.clock + 1
	tm.value = msg
	fmt.Println("ID: ", p.id, " Sending Message: ", msg)
	p.output <- tm
}

func createProcess() *Process {
	p := new(Process)
	p.input = make(chan *Wrapped, max_buf)
	//p.output = make(chan *Wrapped, max_buf)
	p.id = ID
	ID += 1
	return p
}

func (p *Process) run() {
	fmt.Println("RUNNING STUFF", p)
	//go func() {

	//}()
	for i := 0; i < 1; i++ {
		p.sender("RUNNER")
		time.Sleep(time.Millisecond * 1500)
		select {
		case <-p.input:
			fmt.Println("GOT: ", <-p.input)
			Done <- true
		}
	}
	/*select {
	case <-p.input:
		fmt.Println("GOT: ", <-p.input)
		Done <- true
	}*/
	Done <- true
}

func main() {
	Done = make(chan bool, 2)
	p1 := createProcess()
	p2 := createProcess()
	p1.output = p2.input
	p2.output = p1.input
	p1.run()
	go p2.run()
	/*fmt.Println(p)
	t := new(Wrapped)
	t.timestamp = 0
	t.value = "something"
	p.input <- t*/
	/*go func() {
		time.Sleep(time.Second * 1)
		Done <- true
	}()*/
	select {
	case <-Done:
		fmt.Println("Finished")
		/*for n := range p1.input {
			fmt.Println("P1: ", n.value, n.timestamp)
		}*/
		/*for n := range p2.input {
			fmt.Println("P2: ", n.value, n.timestamp)
		}*/
		return
	}
	/*p1 := new(Process)
	p1.input = make(chan *Wrapped)
	//p2 := new(Process)
	t := new(Wrapped)
	t.timestamp = 0
	t.value = "something"
	//c <- t
	fmt.Println(p1)
	p1.input <- t*/
	//time.Sleep(1000)
	//fmt.Println("in p1: ", <-p1.input)
	//p1.receiver()
}
