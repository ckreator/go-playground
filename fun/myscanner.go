package main

import "fmt"
import "regexp"

// create closure-based scanner

func newScanner(s, sep string) func() string{
	ret := regexp.MustCompile(sep).Split(s, -1)
	
	fmt.Println(ret[1])
	i := 0
	return func() string {
		if i < len(ret) {
			for ;ret[i] == "" && i < len(ret); i++ {}
			r := ret[i]
			i += 1
			return r
		}
		return ""
	}
}

func main() {
	f := newScanner("<body>Hello</body>", "<|>")
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
