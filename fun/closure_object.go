package main

import "fmt"

//import "fmt"

type any interface{}

type api struct {
	Set func(key string, value any)
	Get func(key string) any
}

type myObject struct {
	m map[string]any
}

func myObj() *api {
	o := make(map[string]any)
	ret := new(api)
	ret.Set = func(key string, value any) {
		o[key] = value
	}

	ret.Get = func(key string) any {
		return o[key]
	}

	return ret

	//return myObject
}

func main() {
	//myObj(&obj)
	obj := myObj()
	fmt.Println("Obj: ", obj)
	obj.Set("test", 19)
	fmt.Println("obj.test => ", obj.Get("test"))

	//obj("")
}
