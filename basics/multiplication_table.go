package main

import "fmt"
import "strconv"

func multi(count int) {
	// let's just do stupid stuff
	d := 1
	for format := count * count; format > 10; format /= 10 {
		d++
	}
	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			s := "%" + strconv.Itoa(d) + "d "
			fmt.Printf(s, j * i)
		}
		fmt.Println()
	}
}

func main() {
	s := "%" + strconv.Itoa(3)
	fmt.Println(s)
	multi(20)
}
