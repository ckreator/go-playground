package main

import "fmt"

type any interface{}

type Class interface {
	properties map[string]any
	super *Class
}

type Object struct {
	properties map[string]any
	super *Class
}


type SubClassed struct {
	properties map[string]any
	super *Class
}

func newObject() {
	nu := new(Object)
	nu.properties["test"] = 15
	nu.super = &nu
}

// Let's just do this without a constructor
func (obj *Object) constructor() {

}

func main() {
	o := newObject()
}
