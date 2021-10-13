package main

import (
	"fmt"
)

func main() {
	testBase()
}

func testBase() {
	var c1 complex64
	c1 = 3.2 + 12i
	//c2 := 3.2 + 12i
	//c3 := complex(3.2, 12)

	fmt.Println("complex :", real(c1), imag(c1))
}
