package main

import "fmt"

type Animal struct {
	Speed float32
	Weight float32
}

type Dog struct {
	Animal
	Speed float32
}

func main() {
	test()
}

func test() {
	d := &Dog{Animal{100, 200}, 666}
	d.Run()
}

func (animal Animal) Run() float32 {
	fmt.Println("animal speed:", animal.Speed)
	return animal.Speed
}

func (dog Dog) Run() float32 {
	dog.Animal.Run()
	fmt.Println("dog speed:", dog.Speed)
	return dog.Speed
}




