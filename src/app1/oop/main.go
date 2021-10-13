package main

import "fmt"

type Integer int

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	testTypeFunc()
}

func testInterface() {
	var a Integer = 1
	var b LessAdder = &a
	fmt.Println(b)
}

func testTypeFunc() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	a.Add(2)
	fmt.Println(a)
}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}
