package main

import (
	"fmt"
	"strings"
)

func query(s string, m map[string]interface{}) {
	spl := strings.Split(s, ".")
	curr := m
	for _, val := range spl {
		if curr[val] == nil {
			curr[val] = make(map[string]interface{})
		}
		fmt.Println("Curr now: ", curr)
		curr = curr[val].(map[string]interface{})
	}
	fmt.Println(spl)
}

func main() {
	m := make(map[string]interface{})
	m["hello"] = make(map[string]interface{})
	fmt.Println("m: ", m)
	query("a.deeply.nested.object", m)
}
