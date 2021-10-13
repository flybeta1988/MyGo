package main

import (
	"fmt"
	"math"
)

func main() {
	testBase()
}

func testBase() {
	var f32 float32
	f32 = math.MaxFloat32

	foo := 3.14

	fmt.Println("float value range:", f32)
	fmt.Printf("%T\n", foo)
}
