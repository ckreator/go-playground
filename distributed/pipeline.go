package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func extendPipe(doIt func(in <-chan int, out chan int)) func(<-chan int) chan int {
	return func(in <-chan int) chan int {
		out := make(chan int)
		go doIt(in, out)
		return out
	}
}

func add3(in <-chan int, out chan int) {
	for n := range in {
		out <- n + 3
	}
	close(out)
}

func add5(in <-chan int, out chan int) {
	for n := range in {
		out <- n + 5
	}
	close(out)
}

func add7(in <-chan int, out chan int) {
	for n := range in {
		out <- n + 7
	}
	close(out)
}

func main() {
	// Set up the pipeline
	c := gen(2, 3)
	out := square(c)

	/*for n := range out {
		fmt.Println(n)
	}*/

	ary := [3]func(<-chan int, chan int){add3, add5, add7}

	for _, v := range ary {
		out = (extendPipe(v))(out)
	}
	//make([]func(in <-chan int, out chan int), 3)

	//f1 := extendPipe(add3)

	//nu := f1(out)
	for n := range out {
		fmt.Println(n)
	}
}
