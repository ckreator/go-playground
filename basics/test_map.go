package main

import "fmt"

type any interface {}

func main() {
	m := make(map[string]any)
	m["hello"] = "world"
	m["age"] = 5
	fmt.Println(m["hello"])
	fmt.Println(m["age"])
}
