package main

import "fmt"
import "reflect"

func tp(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i.(type))
}

func main() {
	tp("string")
}
