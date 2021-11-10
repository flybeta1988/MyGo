package interview

import "fmt"

type People2 interface {
	Speak(string) string
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

func TestInterface() {
	//var peo People2 = Stduent{}
	var peo People2 = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
