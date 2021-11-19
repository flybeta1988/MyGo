package interview

import "fmt"

type People2 interface {
	Speak(string) string
}

type People3 interface {
	Show()
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func (s *Stduent) Show()  {}

func live() People3 {
	var s *Stduent
	return s
}

func TestInterface() {
	//var peo People2 = Stduent{}
	var peo People2 = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func TestInterface2() {
	result := live()
	fmt.Println(result)
	if result == nil {
		fmt.Println("AAAA")
	} else {
		fmt.Println("BBBB")
	}
}

func TestInterface3() {
	var x *int = nil
	foo(x)
}

func foo(x interface{}) {

	var y interface{}

	fmt.Println(x, y)

	if x == nil {
		fmt.Println("x is empty interface")
	} else {
		fmt.Println("x is non-empty interface")
	}

	if y == nil {
		fmt.Println("y is empty interface")
	} else {
		fmt.Println("y is non-empty interface")
	}
}
