package main

import (
	"fmt"
	"runtime"
	"time"
)

type Machine struct {
	clock int
}

func doMachine(m *Machine, rate int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Clock now at: ", m.clock)
		m.clock += rate
	}
}

func machinesParallel() {
	m1 := new(Machine)
	m2 := new(Machine)
	go doMachine(m1, 3)
	doMachine(m2, 2)
}

func main() {
	runtime.GOMAXPROCS(5)
	m1 := new(Machine)
	m2 := new(Machine)
	doMachine(m1, 3)
	go doMachine(m2, 2)
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("FINISHED MAIN")
}
