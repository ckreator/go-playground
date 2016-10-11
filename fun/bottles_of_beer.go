package main

import (
	"fmt"
	"sync"
	"strconv"
	"time"
)

type locked struct {
	val int
	mux sync.Mutex
}

func bottles(v *locked, amount int, num int) {
	for {
		//fmt.Println("BOTTLES", v.val)
		v.mux.Lock()
		if v.val < 0 {
			v.mux.Unlock();
			break
		}
		fmt.Println(num, " Bottles of beer: ", v.val)
		v.val -= amount
		time.Sleep(500)
		v.mux.Unlock()
	}
}

func doIt(times int) {
	//fmt.Println("BOTTLES", times)
	l := locked{val: 2000000}
	for i := 0; i < times; i++ {
		//fmt.Println("BOTTLES")
		go bottles(&l, 100000, i)
	}
}

func tell(t int, c chan string) {
	//fmt.Println("Say: ", t)
	c <- "Say: " + strconv.Itoa(t)
}

func say(t int) {
	c := make(chan string)
	for i := 0; i < t; i++ {
		go tell(i, c)
		/*select {
		case <- c:
			fmt.Println("Got: ", <-c)
		}*/
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Got: ", <-c)
	}
}

func main() {
	doIt(10)
	//say(3)
}
