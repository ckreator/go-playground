package main

import (
	"fmt"
	"log"
	"strings"
)

type any interface{}

// just a fun little iterator
type InputStream struct {
	next      func() string
	peek      func() string
	eof       func() bool
	suffocate func(string, ...string)
}

func Iterator(input string) *InputStream {
	row := 0
	col := 0
	pos := 0
	ret := new(InputStream)
	splitted := strings.Split(input, "")

	next := func() string {
		c := splitted[pos]
		pos++
		if c == "\n" {
			col = 0
			row++
		} else {
			col++
		}
		return c
	}
	peek := func() string {
		return splitted[pos]
	}
	eof := func() bool {
		return pos >= len(splitted)
	}
	suffocate := func(msg string, list ...string) {
		log.Fatal(msg)
	}

	ret.next = next
	ret.peek = peek
	ret.eof = eof
	ret.suffocate = suffocate
	return ret
}

func main() {
	it := Iterator("hello")
	for !it.eof() {
		fmt.Println(it.next())
	}
}
