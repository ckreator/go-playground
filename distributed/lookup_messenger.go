package main

import (
	"fmt"
	"time"
)

var Table Lookup

type Message struct {
	from      string
	msg       string
	timestamp int
}

type Messenger struct {
	input chan Message
	id    string
}

type Lookup struct {
	_map map[string]*Messenger
}

func (l *Lookup) set(key string, value *Messenger) {
	l._map[key] = value
}

func (l *Lookup) get(key string) *Messenger {
	val, ok := l._map[key]
	if !ok {
		l._map[key] = new(Messenger)
	}
	val = l._map[key]
	if val.input == nil {
		val.input = make(chan Message)
	}
	return val
}

func (m *Messenger) sendTo(id string, msg Message) {
	sender := Table.get(id)
	sender.input <- msg
}

func (m *Messenger) run() {
	//t := 0
	for {
		fmt.Println("In loop")
		select {
		case <-m.input:
			fmt.Println("Got this: ", <-m.input)
			/*ms := new(Message)
			ms.timestamp = t
			t++
			ms.msg = "Sending response"
			m.sendTo(msg.from, *ms)*/
		}
	}
}

func main() {
	m1 := new(Messenger)
	m1.id = "1"
	m1.input = make(chan Message, 2)
	m2 := new(Messenger)
	m2.id = "2"
	m2.input = make(chan Message, 2)
	msg := new(Message)
	msg.from = "1"
	msg.timestamp = 0
	msg.msg = "Hello"
	go m2.run()
	m2.input <- *msg
	//go m1.run()
	time.Sleep(time.Second)
}
