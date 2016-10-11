package main

import "fmt"

// Create a flexible way to create multidimensional arrays

type T interface {
	createMe(int ,int)
}

type myByte uint8

/*func createMe(t T, dy, dx int) [][]T {
	fmt.Print("type: %T\n", t)
	ret := make([][]t, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]t, dx)
	}
	return ret
}*/

func (b myByte) createMe(dy, dx int) [][]myByte {
	fmt.Printf("type: %T\n", b)	
}

func main() {
	var t T
	var b myByte = 9
	t = b
	t.createMe(10, 15)
//	test := createMe(t, 3, 3)
}
